package tracing

import (
	"github.com/go-jedi/auth-service/internal/logger"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

func Init(serviceName string) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		logger.Fatal("failed to init tracing", zap.Error(err))
	}
}
