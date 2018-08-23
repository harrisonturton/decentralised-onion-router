package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Server    ServerConfig `yaml:"server"`
	ExitNode  ExitConfig   `yaml:"exit_node"`
	RelayNode RelayConfig  `yaml:"relay_node"`
}

type ServerConfig struct {
	Host           string  `yaml:"host"`
	Port           string  `yaml:"port"`
	IsTls          bool    `yaml:"tls"`
	TlsKeyFilename *string `yaml:"tls_key"`
	TlsCrtFilename *string `yaml:"tls_crt"`
}

type ExitConfig struct {
	ForceHttps bool `yaml:"force_https"`
	Timeout    int  `yaml:"timeout"`
}

type RelayConfig struct {
	Timeout int `yaml:"timeout"`
}

var DEFAULT_CONFIG = Config{
	Server: ServerConfig{
		Host:           "",
		Port:           "3333",
		IsTls:          false,
		TlsKeyFilename: nil,
		TlsCrtFilename: nil,
	},
	ExitNode: ExitConfig{
		ForceHttps: true,
		Timeout:    10,
	},
	RelayNode: RelayConfig{
		Timeout: 10,
	},
}

func ReadConfig(filename string) (*Config, error) {
	fileBody, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read config file body")
	}

	var config = DEFAULT_CONFIG
	if err := yaml.Unmarshal(fileBody, &config); err != nil {
		return nil, errors.Wrap(err, "Failed to parse config file")
	}

	if config.Server.IsTls && (config.Server.TlsKeyFilename == nil || config.Server.TlsCrtFilename == nil) {
		return nil, errors.Wrap(err, "TLS is enabled but TLS Key or Certificate not supplied.")
	}

	return &config, nil
}
