package Config

import (
  "github.com/BurntSushi/toml"
)

type Base struct {
  Serv Server
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

func Get()(conf Base) {
  _, err := toml.DecodeFile("config/base.toml", &conf)
  if err != nil { panic(err) }
  return conf
}
