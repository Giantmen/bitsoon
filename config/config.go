package config

import (
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Listen   string
	LogLevel string
	LogPath  string
	Debug    bool

	Mysql *Mysql
}

type Mysql struct {
	ConnStr string
	Timeout int
	MaxOpen int
	MaxIdle int
}

func Parse(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	cfg := new(Config)
	err = toml.Unmarshal(buf, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
