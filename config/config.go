package config


import (
	"fmt"
	"github.com/spf13/viper"
)

func InitializeConfig () *viper.Viper {
		v := viper.New()
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		err := v.ReadInConfig()
		if err != nil { 
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		return v
}
