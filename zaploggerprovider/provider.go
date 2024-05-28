package zaploggerprovider

import (
	"fmt"

	"github.com/stateprism/prisma_ca/providers"
	"go.uber.org/zap"
)

type zapLoggerProvider struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

func NewZapLoggerProvider(logger *zap.Logger, sugared *zap.SugaredLogger) *zapLoggerProvider {
	return &zapLoggerProvider{
		logger: logger,
		sugar:  sugared,
	}
}

func (z *zapLoggerProvider) Flush() {
	z.logger.Sync()
}

func (z *zapLoggerProvider) Log(level providers.LogLevel, message string) {
	switch level {
	case providers.LOG_LEVEL_DEBUG:
		if z.logger != nil {
			z.logger.Debug(message)
		} else if z.sugar != nil {
			z.sugar.Debug(message)
		}
	case providers.LOG_LEVEL_INFO:
		if z.logger != nil {
			z.logger.Info(message)
		} else if z.sugar != nil {
			z.sugar.Info(message)
		}
	case providers.LOG_LEVEL_WARN:
		if z.logger != nil {
			z.logger.Warn(message)
		} else if z.sugar != nil {
			z.sugar.Warn(message)
		}
	case providers.LOG_LEVEL_ERROR:
		if z.logger != nil {
			z.logger.Error(message)
		} else if z.sugar != nil {
			z.sugar.Error(message)
		}
	case providers.LOG_LEVEL_CRITICAl:
		if z.logger != nil {
			z.logger.Fatal(message)
		} else if z.sugar != nil {
			z.sugar.Fatal(message)
		}
	default:
		panic("Unknown log level")
	}
}

func (z *zapLoggerProvider) Logf(level providers.LogLevel, format string, args ...interface{}) {
	z.Log(level, fmt.Sprintf(format, args...))
}

func (z *zapLoggerProvider) Fatalf(format string, args ...interface{}) {
	z.Log(providers.LOG_LEVEL_CRITICAl, fmt.Sprintf(format, args...))
}
