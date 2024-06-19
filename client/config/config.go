package config

import "client/services"

type IConfig interface {
	GetService() services.Service
}

type Config struct {
	ApiService services.Service
}

func (c Config) GetService() services.Service {
	return c.ApiService
}

func NewConfig(apiService services.Service) Config {
	return Config{
		ApiService: apiService,
	}
}
