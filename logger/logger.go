package logger

import "fmt"

type LogLevel int

const (
	Debug = iota
	Warning
	Info
	Error
)

func (log *Logger) Log() {

	lg := GetLevelString(log.Level) + ":"
	if log.Level > 0 {
		fmt.Println(lg, log.Message, log.Level)
	}
}

type Logger struct {
	Level   int
	Message string
}

func GetInstance() *Logger {
	log := new(Logger)
	// log.Level = 4
	return log
}

func GetLevelString(logLevel int) string {
	switch logLevel {
	case 0:
		return "Debug"
	case 1:
		return "Warning"
	case 2:
		return "Info"
	case 3:
		return "Error"
	default:
		return "Debug"
	}
}
