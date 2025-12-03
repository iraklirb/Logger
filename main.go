package main

import (
	"logger/logger"
)

func main() {
	log := logger.GetInstance(logger.Info)
	log.Debug("log this message")
	log.Info("log this message")
	log.Warning("log this message")
	log.Error("log this message")
}
