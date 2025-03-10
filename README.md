# cometbft-abci

## run and test

### 1.编译代码
```
go build main.go
```

### 2 使用 cometBFT 初始化 home 目录
```
go run github.com/cometbft/cometbft/cmd/cometbft@v1.0.1 init --home /Users/guoshijiang/CosmosWorkSpace/testdata/.kvstore
```

### 3.启动 CometBft 代码
```
./main -kv-home /Users/guoshijiang/CosmosWorkSpace/testdata/.kvstore
```

### 4.启动一个 Tendermint 节点
```
go run github.com/cometbft/cometbft/cmd/cometbft@v1.0.1 --home /Users/guoshijiang/CosmosWorkSpace/testdata/.kvstore start
```