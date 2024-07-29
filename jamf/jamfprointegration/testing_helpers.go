package jamfprointegration

import "go.uber.org/zap"

func newIntegrationWithLogger() Integration {
	logger, _ := zap.NewProduction()
	return Integration{
		Sugar: logger.Sugar(),
	}
}
