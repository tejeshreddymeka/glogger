// Package config provides configurations for logger
package config

// constants for logger library code(name)
const (
	LOGRUS string = "logrus"
	ZAP    string = "zap"
)

// constants for log level
const (
	DEBUG = "debug"
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
)

// LoggerConfig represents logger configuration
type LoggerConfig struct {
	// logger library code
	Code string `yaml:"code"`
	// log level
	Level string `yaml:"level"`
	// show caller in log message
	EnableCaller bool `yaml:"enableCaller"`
}
