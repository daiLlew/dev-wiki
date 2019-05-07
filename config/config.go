package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Envs struct {
	Develop Env `yaml:"develop"`
}

type Env struct {
	Name        string `yaml:"name"`
	FlorenceURI string `yaml:"florence_uri"`
}

type Config struct {
	Environments Envs `yaml:"environments"`
}

func Get() (*Config, error) {
	b, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
