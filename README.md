
### dubbo-go triple benchmark

before using benchmark, you shuold be sure that triple has import dubbo-go 3.0

derived from grpc-go benchmark


```
bash run_bench.sh --help for help
```

example:
```
bash run_bench.sh -r 1 -c 1 -req 10 -resp 10 -rpc_type unary
```
