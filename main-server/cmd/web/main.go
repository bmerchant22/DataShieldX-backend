package main

import (
	"github.com/bmerchant22/server-controller/pkg/web"
	"go.uber.org/zap"
	"log"
)

func main() {
	var logger *zap.Logger
	var loggerErr error

	if logger, loggerErr = zap.NewDevelopment(); loggerErr != nil {
		log.Fatalln(loggerErr)
	}

	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	web.CreateWebServer()
}
