// Package logrus provides logrus's wrapper function and factory methods
package logrus

import (
	"github.com/pkg/errors"
	"github.com/tejeshreddymeka/glogger"
	"github.com/tejeshreddymeka/glogger/config"
)

// Factory represent logrus factory
type Factory struct{}

// Build method builds the logrus logger as per config
func (l *Factory) Build(loggerConfig *config.LoggerConfig) (glogger.Logger, error) {
	logrusLogger, err := getLogrusLogger(loggerConfig)
	if err != nil {
		return nil, errors.Wrap(err, "getLogrusLogger")
	}
	return logrusLogger, nil
}
