package logger

import (
	"strconv"

	"github.com/fatih/color"
)

type LogLevel int

const (
	Debug = iota
	Warning
	Info
	Error
)

type Logger struct {
	Level   int
	Message string
}

func GetInstance() *Logger {
	log := new(Logger)
	// log.Level = 4
	return log
}

func (log *Logger) Log() {
	lg := GetLevelString(log.Level) + ": " + log.Message + " " + strconv.Itoa(log.Level)
	PrintWithColor(lg, log.Level)
}

func PrintWithColor(logMessage string, level int) {
	gray := color.New(color.FgHiBlack)

	switch level {
	case Debug:
		color.Green(logMessage)
	case Warning:
		color.Yellow(logMessage)
	case Info:
		gray.Println(logMessage)
	case Error:
		color.Red(logMessage)
	default:
		panic("incorrect log level")
	}
}

func GetLevelString(level int) string {
	switch level {
	case Debug:
		return "Debug"
	case Warning:
		return "Warning"
	case Info:
		return "Info"
	case Error:
		return "Error"
	default:
		panic("incorrect log level")
	}
}
