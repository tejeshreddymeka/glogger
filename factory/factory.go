// Package factory provides factory methods of logger
package factory

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tejeshreddymeka/glogger"
	"github.com/tejeshreddymeka/glogger/config"
	"github.com/tejeshreddymeka/glogger/logrus"
	"github.com/tejeshreddymeka/glogger/zap"
)

// loggerFactoryMap maps logger code(name) to specific logger factory
var loggerFactoryMap = map[string]LoggerFactory{
	config.LOGRUS: &logrus.Factory{},
	config.ZAP:    &zap.Factory{},
}

// LoggerFactory represents generic Logger Factory
type LoggerFactory interface {
	Build(*config.LoggerConfig) (glogger.Logger, error)
}

// getLoggerFactory returns logger factory with specified logger code
func getLoggerFactory(loggerCode string) (LoggerFactory, error) {
	loggerFactory, ok := loggerFactoryMap[loggerCode]
	if !ok {
		return nil, fmt.Errorf("Invalid logger code %s", loggerCode)
	}
	return loggerFactory, nil
}

// Build function will build the logger with specified config
func Build(loggerConfig *config.LoggerConfig) (glogger.Logger, error) {
	loggerCode := loggerConfig.Code
	loggerFactory, err := getLoggerFactory(loggerCode)
	if err != nil {
		return nil, errors.Wrap(err, "getLoggerFactory")
	}
	logger, err := loggerFactory.Build(loggerConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "loggerFactory.Build")
	}
	return logger, nil
}
