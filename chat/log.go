package main

import (
	"github.com/Leon2012/xchat2/libs/log"
)

var (
	appLogger *log.Logger
)

func initLogger() error {
	var err error
	logOutput := cfg.Log.Output
	if logOutput == "file" {
		appLogger, err = log.NewFileLogger("routeserver", 0, cfg.Log.File)
	} else if logOutput == "console" {
		appLogger, err = log.NewLogger("routeserver", 0)
	}
	if err != nil {
		return err
	}
	log.SetLevel(log.LEVEL(cfg.Log.Level))
	return nil
}
