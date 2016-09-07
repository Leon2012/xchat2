package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/Leon2012/xchat2/libs/signal"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	flag.Parse()
	if err := initConfig(cfgFile); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	showVersion()
	if err := initLogger(); err != nil {
		fmt.Println("init logger error !")
		os.Exit(-1)
	}
	if err := initRPC(); err != nil {
		fmt.Println("init rpc error !")
		os.Exit(-1)
	}
	appLogger.Info("rpc server %s%s, debug %s start", cfg.Server.Addr, cfg.Server.Rpcpath, cfg.Server.Debugpath)

	signal.InitSignal()
}
