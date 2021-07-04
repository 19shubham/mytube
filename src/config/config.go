// Package config reads a config file in json or yaml and provides
// this to the rest of the packages.
package config

import (
	"fmt"

	"github.com/spf13/viper"

	"constants/enums/environments"
)

// Use this directly across packages
// NOT thread-safe
// Access mutli-level config using . notation
var (
	Environment *environments.Environment
	Config      *viper.Viper
)

// DoInit to initialise application level config data. Config data is kept in a
// settings folder according to the given environment. settings folder in turn
// contains the folder according to the environment(dev, prod etc).
// In the environment folder json files are having the config info.
// Package uses the viper functionality to parse the key value pair
// in json or yaml.
// input params: environment - decided environment of the application machine.

func DoInit(environment string) {
	Config = viper.New()
	Config.SetConfigType("json")
	Config.SetConfigName("config")
	Config.AddConfigPath("settings/" + environment)
	if err := Config.ReadInConfig(); err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//TO Watch config; If not this statement, changes will not reflect.
	Config.WatchConfig()
}
