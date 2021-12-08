package config

import (
	"log"

	"github.com/spf13/viper"
)

/* AppCfg contains all the App configuration loaded from config.yml*/
var AppCfg = loadConfig()

func loadConfig() Config  {
	viper.AddConfigPath(".")
    viper.SetConfigName("app")
    viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8080")
	viper.BindEnv("PORT")
	viper.BindEnv("DATABASE")
	viper.BindEnv("JWTSECRET")

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
    ServerPort string `mapstructure:"PORT"`
    Database string `mapstructure:"DATABASE"`
	JWT_Secret string `mapstructure:"JWTSECRET"`
}
