package main

import (
	"context"
	"fmt"

	framework "github.com/gnolang/gnonative/framework/service"
)

type App struct {
	ctx    context.Context
	bridge *framework.Bridge
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	config := framework.NewBridgeConfig()
	config.UseTcpListener = true
	config.RootDir = "/Users/iuri/dev/tmp"
	config.DisableUdsListener = true
	bridge, err := framework.NewBridge(config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	a.bridge = bridge
	fmt.Println("GnoNative GRPC Server created")
}
