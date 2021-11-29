package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

/* AppCfg contains all the App configuration loaded from config.yml*/
var AppCfg = loadConfig()

func loadConfig() Config  {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
    	log.Fatal(err)
	}
	return cfg
}


type Config struct {
    Server struct {
        Port string `yaml:"port"`
        Host string `yaml:"host"`
    } `yaml:"server"`
    Database struct {
        Uri string `yaml:"uri"`
    } `yaml:"database"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}
