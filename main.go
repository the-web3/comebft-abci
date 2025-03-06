package main

import (
	"fmt"

	abciserver "github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/os"
	db "github.com/tendermint/tm-db"
)

type SimpleApp struct {
	types.BaseApplication
	state db.DB
}

func NewSimpleApp() *SimpleApp {
	state := db.NewMemDB()
	return &SimpleApp{
		state: state,
	}
}

func (app *SimpleApp) CheckTx(req types.RequestCheckTx) types.ResponseCheckTx {
	fmt.Printf("Check tx: %s\n", req.Tx)
	return types.ResponseCheckTx{Code: 0}
}

func (app *SimpleApp) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {
	fmt.Printf("deliver tx: %s\n", req.Tx)

	var tx struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	if err := json.Unmarshal(req.Tx, &tx); err != nil {
		return types.ResponseDeliverTx{Code: 100, Log: "invalid transaction format"}
	}

	err := app.state.Set([]byte(tx.Key), []byte(tx.Value))
	if err != nil {
		return types.ResponseDeliverTx{Code: 100, Log: "set tx fail"}
	}

	fmt.Printf("Transaction Applied: %s=%s\n", tx.Key, tx.Value)

	return types.ResponseDeliverTx{Code: 0}
}

func (app *SimpleApp) Commit() types.ResponseCommit {
	hash, _ := app.state.Get([]byte("hash"))
	return types.ResponseCommit{Data: hash}
}

func main01() {
	app := NewSimpleApp()

	addr := "tcp://127.0.0.1:26658"

	server, err := abciserver.NewServer(addr, "socket", app)
	if err != nil {
		fmt.Println("New server fail", "err", err)
	}

	if err := server.Start(); err != nil {
		fmt.Printf("server start fail")
		os.Exit("1")
	}

	fmt.Printf("ABCI server stated at", addr)

	select {}
}
