package config

import (
	"github.com/joeshaw/envdecode"
)

type config struct {
	Strava stravaAuthConfig
}

type stravaAuthConfig struct {
	ClientId     string   `env:"STRAVA_CLIENT_ID,required"`
	ClientSecret string   `env:"STRAVA_CLIENT_SECRET,required"`
	Scopes       []string `env:"STRAVA_SCOPES,default=read,activity:read_all,activity:write"`
}

func NewConfig() (config, error) {
	cfg := config{}
	err := envdecode.StrictDecode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
