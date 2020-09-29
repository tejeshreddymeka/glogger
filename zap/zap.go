// Package zap provides zap's wrapper function and factory methods
package zap

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/tejeshreddymeka/glogger"
	"github.com/tejeshreddymeka/glogger/config"
	"go.uber.org/zap"
)

// return Zap logger with specified config
func getZapLogger(loggerConfig *config.LoggerConfig) (glogger.Logger, error) {

	logger, err := initZapLogger(loggerConfig)
	if err != nil {
		return nil, errors.Wrap(err, "initLog(loggerConfig)")
	}
	defer logger.Sync()
	zapSugaredlogger := logger.Sugar()

	return zapSugaredlogger, nil
}

// initialize zap logger
func initZapLogger(loggerConfig *config.LoggerConfig) (*zap.Logger, error) {
	// TODO added zap_logger.log as output file, need to configure this as a logger config option
	rawJSONConfig := []byte(`{
	 "level": "info",
     "Development": false,
      "DisableCaller": false,
	 "encoding": "console",
	 "outputPaths": ["stdout", "./zap_logger.log"],
	 "errorOutputPaths": ["stderr"],
	 "encoderConfig": {
		"timeKey":        "timestamp",
		"levelKey":       "level",
		"messageKey":     "message",
        "nameKey":        "name",
		"stacktraceKey":  "stacktrace",
        "callerKey":      "caller",
		"lineEnding":     "\n",
        "timeEncoder":     "ISO8601",
		"levelEncoder":    "lowercase",
        "durationEncoder": "stringDuration",
        "callerEncoder":   "full"
	  }
 	}`)

	var zapConfig zap.Config
	var zapLogger *zap.Logger

	//standard configuration
	err := json.Unmarshal(rawJSONConfig, &zapConfig)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}
	//customize configuration
	err = customizeZapLogger(&zapConfig, loggerConfig)
	if err != nil {
		return nil, errors.Wrap(err, "customizeZapLogger")
	}
	zapLogger, err = zapConfig.Build()
	if err != nil {
		return nil, errors.Wrap(err, "zapConfig.Build")
	}

	return zapLogger, nil
}

// customizeZapLogger customize the zap logger as per specified config
func customizeZapLogger(zapConfig *zap.Config, loggerConfig *config.LoggerConfig) error {
	zapConfig.DisableCaller = !loggerConfig.EnableCaller

	// set log level
	logLevel := zap.NewAtomicLevel().Level()
	err := logLevel.Set(loggerConfig.Level)
	if err != nil {
		return errors.Wrap(err, "logLevel.Set")
	}
	zapConfig.Level.SetLevel(logLevel)

	return nil
}
