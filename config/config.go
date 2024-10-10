package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	configOnce sync.Once
	Config     *config
)

type config struct {
	Server *serverConf `yaml:"server"`
}

type serverConf struct {
	Name string    `yaml:"name"`
	Port *servPort `yaml:"port"`
}

type servPort struct {
	HttpPort string `yaml:"httpPort"`
}

func InitConfig(configFile string) error {
	var err error
	configOnce.Do(func() {
		Config = &config{}
		viper.SetConfigType("yaml")
		viper.SetConfigFile(configFile)
		err = viper.ReadInConfig()
		if err != nil {
			return
		}
	})
	return err
}
