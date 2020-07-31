package logger

import (
	"go.uber.org/zap"
	"strings"
)

type Config struct {
	OutputPath      string
	ErrorOutputPath string
	Level           string
}

func NewLogger(cfg *Config) (*zap.Logger, error) {
	zapCfg := zap.NewProductionConfig()
	if len(cfg.OutputPath) > 0 {
		zapCfg.OutputPaths = []string{cfg.OutputPath}
	}

	if len(cfg.ErrorOutputPath) > 0 {
		zapCfg.ErrorOutputPaths = []string{cfg.ErrorOutputPath}
	}

	if len(cfg.Level) > 0 {
		switch strings.ToLower(cfg.Level) {
		case "debug":
			zapCfg.Level.SetLevel(zap.DebugLevel)
		case "info":
			zapCfg.Level.SetLevel(zap.InfoLevel)
		case "warn":
			zapCfg.Level.SetLevel(zap.WarnLevel)
		case "error":
			zapCfg.Level.SetLevel(zap.ErrorLevel)
		}
	}

	zapCfg.Encoding = "console"

	return zapCfg.Build()
}
