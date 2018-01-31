// Package logger
// 3 January, 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package logger

import (
	"io"
	"log"
)

var logLevel int

// Logger struct
type Logger struct {
	Debug Level

	Info Level

	Notice Level

	Error Level
}

// Level contains the logger level
type Level struct {
	*log.Logger
}

// Init initializes the logger
func (l *Logger) Init(debugHandle io.Writer, infoHandle io.Writer, noticeHandle io.Writer, errorHandle io.Writer, level int) {
	l.Debug.Logger = log.New(debugHandle, "DEBUG: ", log.Ldate|log.Ltime)
	l.Info.Logger = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime)
	l.Notice.Logger = log.New(noticeHandle, "NOTICE: ", log.Ldate|log.Ltime)
	l.Error.Logger = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime)
	logLevel = level
}

// Log writes a message to the log
func (l *Level) Log(format string, a ...interface{}) {
	switch l.Logger.Prefix() {
	case "DEBUG: ":
		if logLevel == 0 {
			l.Logger.Printf(format, a...)
		}
	case "INFO: ":
		if logLevel <= 1 {
			l.Logger.Printf(format, a...)
		}
	case "NOTICE: ":
		if logLevel <= 2 {
			l.Logger.Printf(format, a...)
		}
	case "ERROR: ":
		if logLevel <= 3 {
			l.Logger.Printf(format, a...)
		}
	}
}
