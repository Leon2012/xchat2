package main

/*
 * Created on Fri Sep 09 2016
 *
 * Copyright (c) 2016 Leon.peng@live.com
 */

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/Leon2012/xchat2/libs/idgen"
	"github.com/Leon2012/xchat2/libs/signal"
)

var buildstamp = ""
var globals struct {
	sessionStore *SessionStore
	idGen        *idgen.IdGenerator
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

	globals.idGen = &idgen.IdGenerator{}
	globals.idGen.Init(uint(cfg.Server.Id), []byte(cfg.Server.Encryptionkey))

	globals.sessionStore = newSessionStore(15 * time.Second)
	ws_init(cfg.Server.Name, cfg.Server.Addr)

	appLogger.Info("websocket server start %s ", cfg.Server.Addr+"/chat")

	signal.InitSignal()
}
