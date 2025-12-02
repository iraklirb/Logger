package main

import (
	"logger/logger"
)

func main() {
	// log := logger.Logger{Level: 1, Message: "this message came from Level:"}
	log := logger.GetInstance()
	log.Message = "new logger created"

	log.Level = 1
	log.Log()
	log.Level = 2
	log.Log()
	log.Level = 3
	log.Log()
	log.Level = 0
	log.Log()

	// logger.Log(logger.Logger{Level: 1, Message: "this message came from Level:"})
}
