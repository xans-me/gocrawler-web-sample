package logger

import (
	"github.com/sirupsen/logrus"
	"gocrawler-web-sample/infrastructure/environment"
)

// loggerOptions models struct
type loggerOptions struct {
	basePath         string
	productionPrefix string
	devPrefix        string
	devLevel         logrus.Level
	stagingLevel     logrus.Level
	productionLevel  logrus.Level
	fileTemplate     string
}

func (options *loggerOptions) getPrefix(env environment.AppEnvironment) string {
	if env.IsDev() {
		return options.devPrefix
	} else {
		return options.productionPrefix
	}
}

func (options *loggerOptions) getLevel(env environment.AppEnvironment) logrus.Level {
	if env.IsDev() {
		return options.devLevel
	} else {
		return options.productionLevel
	}
}

// loggerOption models
type loggerOption func(*loggerOptions)

func BasePath(path string) loggerOption {
	return func(options *loggerOptions) {
		options.basePath = path
	}
}

func DevelopmentPrefix(prefix string) loggerOption {
	return func(options *loggerOptions) {
		options.devPrefix = prefix
	}
}

// ProductionPrefix func
func ProductionPrefix(prefix string) loggerOption {
	return func(options *loggerOptions) {
		options.productionPrefix = prefix
	}
}

func FileTemplate(template string) loggerOption {
	return func(options *loggerOptions) {
		options.fileTemplate = template
	}
}

func DevelopmentLevel(level logrus.Level) loggerOption {
	return func(options *loggerOptions) {
		options.devLevel = level
	}
}

func StagingLevel(level logrus.Level) loggerOption {
	return func(options *loggerOptions) {
		options.stagingLevel = level
	}
}

func ProductionLevel(level logrus.Level) loggerOption {
	return func(options *loggerOptions) {
		options.productionLevel = level
	}
}

func (options *loggerOptions) apply(setters []loggerOption) {
	for _, setter := range setters {
		setter(options)
	}
}

func createDefaultOptions() *loggerOptions {
	return &loggerOptions{
		basePath:         "./logs",
		productionPrefix: "prod",
		devPrefix:        "dev",
		fileTemplate:     "app_%Y-%m-%d-%H%M.log",
		devLevel:         logrus.DebugLevel,
		stagingLevel:     logrus.InfoLevel,
		productionLevel:  logrus.InfoLevel,
	}
}
