package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Promcfg struct {
	Server	string `yaml:"Server"`
	Port    int    `yaml:"Port"`
	Intval  int    `yaml:"Intval"`
}

type Config struct {
	PromPushGW Promcfg `yaml:"PromPushGW"`
}

var Monitor Config

func Load(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("configure file doesn't exist")
		os.Exit(-1)
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("read configure file fail, %s\n", err.Error())
		os.Exit(-1)
	}

	err = yaml.Unmarshal(content, &Monitor)
	if err != nil {
		fmt.Printf("configure file format error, %s\n", err.Error())
		os.Exit(-1)
	}
}
