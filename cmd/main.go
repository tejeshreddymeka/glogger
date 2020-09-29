// Package main provides entrylevel code
package main

import (
	"log"

	"github.com/pkg/errors"
	"github.com/tejeshreddymeka/glogger/config"
	"github.com/tejeshreddymeka/glogger/factory"
)

func main() {
	loggerConfig := config.LoggerConfig{
		Code:         config.LOGRUS,
		Level:        config.DEBUG,
		EnableCaller: true,
	}

	logger, err := factory.Build(&loggerConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "factory.Build"))
	}

	logger.Infof("Info %s ", "..!")
	logger.Debugf("Debug ..!")
	logger.Warnf("Warn ..!")

	loggerConfig.EnableCaller = false
	logger, err = factory.Build(&loggerConfig)
	logger.Infof("Info %s ", "..!")
	logger.Debugf("Debug ..!")
	logger.Warnf("Warn ..!")

	loggerConfig.Code = config.ZAP
	loggerConfig.EnableCaller = true
	logger, err = factory.Build(&loggerConfig)
	logger.Infof("Info %s ", "..!")
	logger.Debugf("Debug ..!")
	logger.Warnf("Warn ..!")

	loggerConfig.EnableCaller = false
	logger, err = factory.Build(&loggerConfig)
	logger.Infof("Info %s ", "..!")
	logger.Debugf("Debug ..!")
	logger.Warnf("Warn ..!")
}
