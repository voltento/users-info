package logger

import "go.uber.org/zap"

func NewMock() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.ErrorOutputPaths, cfg.OutputPaths = []string{}, []string{}
	cfg.Level.SetLevel(zap.FatalLevel)
	logger, err := cfg.Build()
	if err != nil {
		panic(err.Error())
	}
	return logger
}
