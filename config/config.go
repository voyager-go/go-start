package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var (
	ConfEnv string
	Cfg     *Conf
	AppPort = "8080"
)

type Conf struct {
	Server ServerConf
	Log    LogConf
}

type ServerConf struct {
	Mode string
	Port string
}

type LogConf struct {
	Debug    bool
	FileName string `yaml:"fileName"`
	DirPath  string `yaml:"dirPath"`
}

func NewConfig() {
	var (
		configPath = fmt.Sprintf("./config/config.%s.yaml", ConfEnv)
	)
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		log.Fatalln(err)
	}
}
