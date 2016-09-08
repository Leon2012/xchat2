package main

import (
	"errors"
	"flag"
	"strconv"
	"strings"

	"github.com/Leon2012/xchat2/libs/config"
	"github.com/Leon2012/xchat2/libs/file"
)

type Config struct {
	Server struct {
		Id      int
		Addr   string
	}
	Log struct {
		Output string
		Level  int32
		File   string
	}
	Rpc struct {
		Addr string
	}
}

var (
	cfg     *Config
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "c", "./chat.ini", "set config file")
}

func initConfig(cfgFile string) error {
	if cfgFile == "" {
		return errors.New("config file empty!")
	}
	exist, _ := file.FileExists(cfgFile)
	if !exist {
		return errors.New("config file not exist!")
	}
	cfg = &Config{}
	err := config.LoadConfigFromFile(cfgFile, cfg)
	if err != nil {
		return err
	}
	return nil
}