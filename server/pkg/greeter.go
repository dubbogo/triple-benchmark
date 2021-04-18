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

package pkg

import (
	"context"
	"fmt"
	pb "github.com/dubbogo/triple/benchmark/protobuf"
)

type GreeterProvider struct {
	*pb.Dubbo3GreeterProviderBase
}

func NewGreeterProvider() *GreeterProvider {
	return &GreeterProvider{
		Dubbo3GreeterProviderBase: &pb.Dubbo3GreeterProviderBase{},
	}
}

func (s *GreeterProvider) BigStreamTest(svr pb.Dubbo3Greeter_BigStreamTestServer) error {
	c, err := svr.Recv()
	if err != nil {
		return err
	}
	fmt.Println("server server recv 1 = ", len(c.Data))

	err = svr.Send(&pb.BigData{
		Data:     make([]byte, c.WantSize),
		WantSize: 0,
	})

	if err != nil {
		fmt.Println("server Send error", len(c.Data))
	}

	return nil
}

// Dubbo3SayHello2 is a unary-client rpc example
func (s *GreeterProvider) Dubbo3SayHello2(ctx context.Context, in *pb.Dubbo3HelloRequest) (*pb.Dubbo3HelloReply, error) {
	fmt.Println("######### get server request name :" + in.Myname)
	fmt.Println("get tri-req-id = ", ctx.Value("tri-req-id"))
	return &pb.Dubbo3HelloReply{Msg: "Hello " + in.Myname}, nil
}

// Dubbo3SayHello is a server rpc exmple
func (g *GreeterProvider) Dubbo3SayHello(svr pb.Dubbo3Greeter_Dubbo3SayHelloServer) error {
	c, err := svr.Recv()
	if err != nil {
		return err
	}
	fmt.Println("server server recv 1 = ", c)
	c2, err := svr.Recv()
	if err != nil {
		return err
	}
	fmt.Println("server server recv 2 = ", c2)
	c3, err := svr.Recv()
	if err != nil {
		return err
	}
	fmt.Println("server server recv 3 = ", c3)

	if err := svr.Send(&pb.Dubbo3HelloReply{
		Msg: c.Myname + c2.Myname,
	}); err != nil {
		panic(err)
	}
	fmt.Println("server server send 1 = ", c.Myname+c2.Myname)
	if err := svr.Send(&pb.Dubbo3HelloReply{
		Msg: c3.Myname,
	}); err != nil {
		panic(err)
	}
	fmt.Println("server server send 2 = ", c3.Myname)
	return nil
}
func (s *GreeterProvider) BigUnaryTest(ctx context.Context, in *pb.BigData) (*pb.BigData, error) {
	fmt.Println("server unary recv len = ", len(in.Data))
	return &pb.BigData{
		Data: make([]byte, in.WantSize),
	}, nil
}

func (g *GreeterProvider) Reference() string {
	return "GrpcGreeterImpl"
}
