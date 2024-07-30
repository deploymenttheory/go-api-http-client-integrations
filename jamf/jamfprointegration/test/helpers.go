package test

import "go.uber.org/zap"

func newSugaredDevelopmentLogger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}
