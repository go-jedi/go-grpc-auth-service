package config

import (
	"os"

	"github.com/pkg/errors"
)

const (
	serviceNameEnvName = "SERVICE_NAME"
)

type ServiceNameConfig interface {
	Service() string
}

type serviceNameConfig struct {
	service string
}

func NewServiceNameConfig() (ServiceNameConfig, error) {
	service := os.Getenv(serviceNameEnvName)
	if len(service) == 0 {
		return nil, errors.New("service name not found")
	}

	return &serviceNameConfig{
		service: service,
	}, nil
}

func (cfg *serviceNameConfig) Service() string {
	return cfg.service
}
