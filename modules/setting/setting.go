// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package setting

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// EnvConfig defines all available environment variables.
type EnvConfig struct {
	Address   string `envconfig:"API_ADDRESS" default:"0.0.0.0:50000"`
	DbPath    string `envconfig:"DB_PATH" default:"gpnm.db"`
	JwtSecret string `envconfig:"JWT_SECRET" default:"gpnm jwt secret"`
}

const (
	AccessToken = "access_token"
)

var (
	Config EnvConfig
)

func init() {
	err := envconfig.Process("", &Config)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
