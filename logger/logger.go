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
	MaxLevel
)

type Logger struct {
	level LogLevel
}

func GetInstance(logLevel LogLevel) *Logger {
	if logLevel < 0 || logLevel >= MaxLevel {
		panic("incorrect log level")
	}
	log := new(Logger)
	log.level = logLevel
	return log
}

func (log *Logger) Debug(message string) {
	if log.level > Debug {
		return
	}
	lg := "Debug: " + message + strconv.Itoa(int(Debug))
	PrintWithColor(lg, Debug)
}

func (log *Logger) Warning(message string) {
	if log.level > Warning {
		return
	}
	lg := "Warning: " + message + strconv.Itoa(int(Warning))
	PrintWithColor(lg, Warning)
}

// if (a || b) && c

func (log *Logger) Info(message string) {
	if log.level > Info {
		return
	}
	lg := "Info: " + message + strconv.Itoa(int(Info))
	PrintWithColor(lg, Info)
}

func (log *Logger) Error(message string) {
	if log.level > Error {
		return
	}
	lg := "Error: " + message + strconv.Itoa(int(Error))
	PrintWithColor(lg, Error)
}

func PrintWithColor(logMessage string, level LogLevel) {
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
