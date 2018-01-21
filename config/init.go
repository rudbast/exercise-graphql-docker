package config

import (
	"errors"
	"flag"
	"fmt"

	gcfg "gopkg.in/gcfg.v1"
)

type (
	Configuration struct {
		Debug    bool
		Server   serverConfig
		Database databaseConfig
	}

	serverConfig struct {
		Host    string
		Timeout string
	}

	databaseConfig struct {
		Host string
		User string
		Pass string
		Name string
	}
)

var (
	// Singleton module.
	globalCfg Configuration

	// Config file name (full path) parsed from command parameter.
	configFilename string
	debug          bool

	// Custom errors.
	ErrMissingConfigFile = errors.New("missing valid config file")
)

func init() {
	// Config file path parameter.
	flag.StringVar(&configFilename, "config", "", "path to config file")
	flag.BoolVar(&debug, "debug", false, "application debugging mode")
}

func Init() error {
	if configFilename == "" {
		return ErrMissingConfigFile
	}

	// Try to read given config file if exists.
	err := gcfg.ReadFileInto(&globalCfg, configFilename)
	if err != nil {
		return err
	}

	globalCfg.Debug = debug
	return nil
}

func Get() Configuration {
	return globalCfg
}

func (dc databaseConfig) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		dc.User, dc.Pass, dc.Host, dc.Name)
}
