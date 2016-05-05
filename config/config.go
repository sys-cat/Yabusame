package Config

import (
  "github.com/BurntSushi/toml"
)

type Base struct {
  Elasticsearch Elasticsearch
  Server Server
  Auth Auth
}

type Server struct {
  PORT string `toml:"PORT"`
  ENV string `toml:"ENV"`
}

type Auth struct {
  SEC string `toml:"SECRET"`
  SALT string `toml:"SALT"`
}

type Elasticsearch struct {
  URL string `toml:"URL"`
  CLIENT_PORT string `toml:"CLIENT_PORT"`
  BASE_PORT string `toml:"BASE_PORT"`
}

func Get()(conf Base) {
  _, err := toml.DecodeFile("config/base.toml", &conf)
  if err != nil { panic(err) }
  return conf
}
