package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port int    `mapstructure:"PORT"`
	Env  string `mapstructure:"ENVIRONMENT"`
}

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return 
}
