package internal

import (
	"log"

	"go.uber.org/zap"
)

func init() {
	var config = zap.NewDevelopmentConfig()
	_logger, err := config.Build()
	if err != nil {
		log.Fatal((err))
	}
	logger = *_logger.Sugar()
}

var logger zap.SugaredLogger

func Logger() zap.SugaredLogger {
	return logger
}
