/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/dubbo-go/common/logger"
	"github.com/dubbogo/triple/benchmark/client/pkg"
	pb "github.com/dubbogo/triple/benchmark/protobuf"
	"github.com/dubbogo/triple/benchmark/stats"
	"github.com/dubbogo/triple/internal/syscall"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	"github.com/apache/dubbo-go/config"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/protocol/dubbo3"
	_ "github.com/apache/dubbo-go/protocol/grpc"
	_ "github.com/apache/dubbo-go/registry/protocol"
	_ "github.com/apache/dubbo-go/registry/zookeeper"
)

var (
	numRPC    = flag.Int("r", 10, "The number of concurrent RPCs on each connection.")
	numConn   = flag.Int("c", 10, "The number of parallel connections.")
	warmupDur = flag.Int("w", 10, "Warm-up duration in seconds")
	duration  = flag.Int("d", 60, "Benchmark duration in seconds")
	rqSize    = flag.Int("req", 30, "Request message size in bytes.")
	rspSize   = flag.Int("resp", 30, "Response message size in bytes.")
	rpcType   = flag.String("streaming", "unary",
		`Configure different client rpc type. Valid options are:
		   unary;
		   streaming.`)
	testName = flag.String("test_name", "client", "Name of the test used for creating profiles.")
	wg       sync.WaitGroup
	hopts    = stats.HistogramOptions{
		NumBuckets:   2495,
		GrowthFactor: .01,
	}
	mu    sync.Mutex
	hists []*stats.Histogram
)

var grpcGreeterImpl = new(pkg.GrpcGreeterImpl)

func init() {
	config.SetConsumerService(grpcGreeterImpl)
}

func main() {

	config.Load()
	time.Sleep(3 * time.Second)

	BigDataReq := pb.BigData{
		WantSize: int32(*rspSize),
		Data:     make([]byte, *rqSize),
	}

	warmDeadline := time.Now().Add(time.Duration(*warmupDur) * time.Second)
	endDeadline := warmDeadline.Add(time.Duration(*duration) * time.Second)
	cf, err := os.Create("/tmp/" + *testName + ".cpu")
	if err != nil {
		logger.Error("Error creating file: %v", err)
	}
	defer func() {
		if err := cf.Close(); err != nil {
			panic(err)
		}
	}()

	if err := pprof.StartCPUProfile(cf); err != nil {
		panic(err)
	}
	cpuBeg := syscall.GetCPUTime()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "tri-req-id", "test_value_XXXXXXXX")
	runWithClient(ctx, &BigDataReq, warmDeadline, endDeadline)

	wg.Wait()
	fmt.Println("handle mem and cpu")
	cpu := time.Duration(syscall.GetCPUTime() - cpuBeg)
	pprof.StopCPUProfile()
	mf, err := os.Create("/tmp/" + *testName + ".mem")
	if err != nil {
		logger.Error("Error creating file: %v", err)
	}
	defer func() {
		if err := mf.Close(); err != nil {
			panic(err)
		}
	}()
	runtime.GC() // materialize all statistics
	if err := pprof.WriteHeapProfile(mf); err != nil {
		logger.Error("Error writing memory profile: %v", err)
	}
	hist := stats.NewHistogram(hopts)
	for _, h := range hists {
		hist.Merge(h)
	}
	parseHist(hist)
	fmt.Println("Client CPU utilization:", cpu)
	fmt.Println("Client CPU profile:", cf.Name())
	fmt.Println("Client Mem Profile:", mf.Name())

}

func runWithClient(ctx context.Context, in *pb.BigData, warmDeadline, endDeadline time.Time) {

	for i := 0; i < *numConn; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			caller := makeCaller(in)
			hist := stats.NewHistogram(hopts)
			for {
				start := time.Now()
				if start.After(endDeadline) {
					mu.Lock()
					hists = append(hists, hist)
					mu.Unlock()
					logger.Info("finish one goroutine")
					return
				}

				if *rpcType == "unary" {
					caller()
				} else {
					for j := 0; j < *numRPC; j++ {
						caller()
					}
				}

				elapsed := time.Since(start)
				if start.After(warmDeadline) {
					if err := hist.Add(elapsed.Nanoseconds()); err != nil {
						panic(err)
					}
				}
			}
		}()
	}
}

func makeCaller(in *pb.BigData) func() {
	if *rpcType == "unary" {
		return func() {
			logger.Info("call RPC")
			if _, err := grpcGreeterImpl.BigUnaryTest(context.Background(), in); err != nil {
				logger.Info("RPC failed: %v", err)
			} else {
				logger.Info("RPC success")
			}
		}
	}

	stream, err := grpcGreeterImpl.BigStreamTest(context.Background())
	if err != nil {
		logger.Errorf("RPC failed: %v", err)
	}
	return func() {

		if err := stream.Send(in); err != nil {
			logger.Errorf("Streaming RPC failed to send: %v", err)
		}
		if _, err := stream.Recv(); err != nil {
			logger.Errorf("Streaming RPC failed to read: %v", err)
		}
	}
}

func parseHist(hist *stats.Histogram) {
	fmt.Println("qps:", float64(hist.Count)/float64(*duration))
	fmt.Printf("Latency: (50/90/99 %%ile): %v/%v/%v\n",
		time.Duration(median(.5, hist)),
		time.Duration(median(.9, hist)),
		time.Duration(median(.99, hist)))
}

func median(percentile float64, h *stats.Histogram) int64 {
	need := int64(float64(h.Count) * percentile)
	have := int64(0)
	for _, bucket := range h.Buckets {
		count := bucket.Count
		if have+count >= need {
			percent := float64(need-have) / float64(count)
			return int64((1.0-percent)*bucket.LowBound + percent*bucket.LowBound*(1.0+hopts.GrowthFactor))
		}
		have += bucket.Count
	}
	panic("should have found a bound")
}
