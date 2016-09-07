package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	APP_VERSION = "0.0.1"
)

var version bool

func init() {
	flag.BoolVar(&version, "version", false, "show version")
}

func showVersion() {
	if version == true {
		fmt.Printf("version : %s \n", APP_VERSION)
		os.Exit(0)
	}
}
