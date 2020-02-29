package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var Conf Config

type Config struct {
	Redis   `toml:"redis"`
	DB      `toml:"database"`
	General `toml:"general"`
	Logger  `toml:"logger"`
}

type Redis struct {
	Addr     string
	Password string
}

type DB struct {
	Addr         string
	User         string
	Password     string
	DatabaseName string
}

type General struct {
	Hostname           string
	RecaptchaSecretKey string
	Listen             string
	Port               int
}

type Logger struct {
	Normal string
	Error  string
}

func (c *Config) InitConfFromToml(path string) {
	if _, err := toml.DecodeFile(path, &Conf); err != nil {
		log.Fatal("Configuration init error")
	}
}
