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
		Addr      string
		Rpcpath   string
		Debugpath string
	}
	Log struct {
		Output string
		Level  int32
		File   string
	}
	Common struct {
		Cpuprofile string
	}

	Message struct {
		Addrs   []string
		Servers []*MsgServer
	}
}

type MsgServer struct {
	sid  int
	addr string
}

var (
	cfg     *Config
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "c", "./route.ini", "set config file")
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
	for _, addr := range cfg.Message.Addrs {
		info := strings.Split(addr, "|")
		sid, err := strconv.Atoi(info[0])
		saddr := info[1]
		if err == nil {
			msgServer := &MsgServer{
				sid:  sid,
				addr: saddr,
			}
			cfg.Message.Servers = append(cfg.Message.Servers, msgServer)
		}
	}
	return nil
}
