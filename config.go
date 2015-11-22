package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Message string `toml:"message"`
}

func NewConfig(path string) *Config {
	var conf Config
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		log.Fatalf("cannot read %v: %v", path, err)
	}

	return &conf
}
