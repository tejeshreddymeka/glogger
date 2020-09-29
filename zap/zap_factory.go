// Package zap provides zap's wrapper function and factory methods
package zap

import (
	"github.com/pkg/errors"
	"github.com/tejeshreddymeka/glogger"
	"github.com/tejeshreddymeka/glogger/config"
)

// Factory represent Zap factory
type Factory struct{}

// Build method will build Zap logger as per config
func (z *Factory) Build(loggerConfig *config.LoggerConfig) (glogger.Logger, error) {
	zapLogger, err := getZapLogger(loggerConfig)
	if err != nil {
		return zapLogger, errors.Wrap(err, "getZapLogger")
	}
	return zapLogger, nil
}
