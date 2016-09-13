package main

import (
	"errors"
	"flag"
	"strings"

	"github.com/Leon2012/xchat2/libs/config"
	"github.com/Leon2012/xchat2/libs/file"
	"github.com/Leon2012/xchat2/store/db/mongo"
)

type Config struct {
	Server struct {
		Id   int
		Name string
		Addr string
	}
	Log struct {
		Output string
		Level  int32
		File   string
	}
	Route struct {
		Addrs   []string
		Servers []*Server
	}
	Dbconf mongo.MongoConfig
}

type Server struct {
	Name string
	Addr string
}

var (
	cfg     *Config
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "c", "./logic.ini", "set config file")
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
	cfg.Route.Servers = parseServerAddrs(cfg.Route.Addrs)
	return nil
}

func reloadConfig() (*Config, error) {
	newCfg := &Config{}
	err := config.LoadConfigFromFile(cfgFile, newCfg)
	if err != nil {
		return nil, err
	}
	newCfg.Route.Servers = parseServerAddrs(newCfg.Route.Addrs)
	cfg = newCfg
	return newCfg, nil
}

func parseServerAddrs(addrs []string) []*Server {
	var servers []*Server
	for _, addr := range addrs {
		info := strings.Split(addr, "|")
		sname := info[0]
		saddr := info[1]
		server := &Server{
			Name: sname,
			Addr: saddr,
		}
		servers = append(servers, server)
	}
	return servers
}
