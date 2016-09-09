package main

import (
	"errors"
	"flag"
	"strings"

	"github.com/Leon2012/xchat2/libs/config"
	"github.com/Leon2012/xchat2/libs/file"
)

type Config struct {
	Server struct {
		Id            int
		Name          string
		Addr          string
		Encryptionkey string
	}
	Log struct {
		Output string
		Level  int32
		File   string
	}
	Rpc struct {
		Addr string
	}
	Logic struct {
		Addrs   []string
		Servers []*Server
	}
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
	cfg.Logic.Servers = parseServerAddrs(cfg.Logic.Addrs)
	return nil
}

func reloadConfig() (*Config, error) {
	newCfg := &Config{}
	err := config.LoadConfigFromFile(cfgFile, newCfg)
	if err != nil {
		return nil, err
	}
	newCfg.Logic.Servers = parseServerAddrs(newCfg.Logic.Addrs)
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
