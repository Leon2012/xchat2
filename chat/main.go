package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/Leon2012/xchat2/libs/signal"
)

var buildstamp = ""
var globals struct {
	sessionStore *SessionStore
}

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

	globals.sessionStore = newSessionStore(15 * time.Second)
	ws_init("", cfg.Server.Addr)

	signal.InitSignal()
}
