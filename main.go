package main

import (
	"logger/logger"
)

func main() {
	// log := logger.Logger{Level: 1, Message: "this message came from Level:"}
	log := logger.GetInstance()
	log.Level = 3
	log.Message = "new logger created"
	log.Log()
	// logger.Log(logger.Logger{Level: 1, Message: "this message came from Level:"})
}
