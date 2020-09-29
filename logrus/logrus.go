// Package logrus provides logrus's wrapper function and factory methods
package logrus

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tejeshreddymeka/glogger"
	"github.com/tejeshreddymeka/glogger/config"
)

// getLogrusLogger return Logrus logger with specified configurations
func getLogrusLogger(loggerConfig *config.LoggerConfig) (glogger.Logger, error) {
	logger := logrus.New()
	//standard configuration
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetReportCaller(true)

	//customize it as per configurations
	err := customizeLogrusLogger(logger, loggerConfig)
	if err != nil {
		return nil, errors.Wrap(err, "customizeLogrusLogger")
	}
	return logger, nil
}

// customizeLogrusLogger customize the logrus logger as per specified configurations
func customizeLogrusLogger(logger *logrus.Logger, loggerConfig *config.LoggerConfig) error {
	logger.SetReportCaller(loggerConfig.EnableCaller)
	logLevel := &logger.Level
	err := logLevel.UnmarshalText([]byte(loggerConfig.Level))
	if err != nil {
		return errors.Wrap(err, "logLevel.UnmarshalText")
	}
	logger.SetLevel(*logLevel)
	return nil
}
