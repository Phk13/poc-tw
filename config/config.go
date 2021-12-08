package config

import (
	"log"

	"github.com/spf13/viper"
)

/* AppCfg contains all the App configuration loaded from config.yml*/
var AppCfg = loadConfig()

func loadConfig() Config  {
	viper.AddConfigPath(".")
    viper.SetConfigName("config")
    viper.SetConfigType("yml")
	viper.AutomaticEnv()
	viper.SetDefault("ServerPort", "8080")

	var cfg Config
	
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}
	return cfg
}


type Config struct {
    ServerPort string `mapstructure:"serverport"`
    Database string `mapstructure:"database"`
	JWT_Secret string `yaml:"jwt_secret"`
}
