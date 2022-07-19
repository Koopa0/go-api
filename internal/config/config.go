package config

import "log"

type AppConfig struct {
	UseCache     bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	JwtConfig    JwtConfig
}

type JwtConfig struct {
	env string
	jwt struct {
		secret string
	}
}
