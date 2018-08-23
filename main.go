package main

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"onion-router/server"
	"os"
)

type YamlConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

const CONFIG_FILENAME = "config.yaml"

func main() {
	fmt.Println("Starting decentralised onion router...")

	config, err := ReadConfig(CONFIG_FILENAME)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Could not read config file").Error())
		os.Exit(1)
	}
	server.Serve(*config)
}

func ReadConfig(filename string) (*server.Config, error) {
	fileBody, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read config file body")
	}

	var config server.Config
	if err := yaml.Unmarshal(fileBody, &config); err != nil {
		return nil, errors.Wrap(err, "Failed to parse config file")
	}

	return &config, nil
}
