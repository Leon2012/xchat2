package config

import (
	"code.google.com/p/gcfg"
)

func LoadConfigFromFile(cfgFile string, cfgObj interface{}) error {
	var err error
	err = gcfg.ReadFileInto(cfgObj, cfgFile)
	if err != nil {
		return err
	}
	return nil
}
