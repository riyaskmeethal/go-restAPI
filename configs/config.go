package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Configs struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

func ReadConf() (cfg Configs) {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Println(err)
	}
	return
}
