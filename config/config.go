package config


import (
	"os"
	"gopkg.in/yaml.v2"
	"log"
)


type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}


var Cfg Config


func ReadConfig() {
	f, err := os.Open("config/config.yml")
	if err != nil {
		log.Println("err: ", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		log.Println("err: ", err)
	}

	log.Println("config: ", Cfg)

}

func GetConfig() *Config {
	return &Cfg
}