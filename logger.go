// Package logger
// 3 January, 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package logger

import (
	"os"

	"github.com/logrusorgru/aurora"
)

// The LogLevel type represents the current logging level of the logger.
type LogLevel uint8

// The ColorFormat type represents formatting flags for the colorizer.
type ColorFormat uint8

// Constants for defining LogLevels.
const (
	All        LogLevel = iota // All events will be logged.
	Verbose                    // Debug events will not be shown.
	Normal                     // Info events will not be shown.
	ErrorsOnly                 // Notice events will not be shown.
	Test                       // No events will be shown.
)

// Timestamp format flags can be used in the format (dateflag|hourflag|timezoneflag).
// If you try to set or use an invalid set of flags (two dateflags or two hourflags) you cause an error.
const (
	ShortDate  = 1 << iota //"1/2/2006"
	LongDate               //"2 Jan 2006"
	Time12Hour             //"3:04:05 PM"
	Time24Hour             //"15:05:05"
	TimeZone               //"MST"
	datemask   = ShortDate | LongDate
	hourmask   = Time12Hour | Time24Hour
	timemask   = TimeZone
)

// A Logger represents a collection of event loggers.
type Logger struct {
	logLevel  LogLevel
	timestamp bool
	colored   bool
	au        aurora.Aurora
	toDisk    bool
	logPath   string
	Debug     Event // Debug event controller
	Info      Event // Info event controller
	Notice    Event // Notice event controller
	Error     Event // Error event controller
}

// Color format flags for determining which parts of an event log get colored.
// If either the Logger's or corresponding Event's colored flag is flase, no colors will be displayed. The '-' in the log will not be colored.
const (
	Timestamp ColorFormat = 1 << iota
	Prefix
	Message
	cformatMask = Timestamp | Prefix | Message
)

// Wrappers for aurora special formats.
const (
	Bold aurora.Color = 1 << iota
	Inverse
)

// Wrappers for aurora foreground colors.
const (
	BlackFg aurora.Color = (1 + iota) << 8
	RedFg
	GreenFg
	YellowFg
	BlueFg
	MagentaFg
	CyanFg
	GrayFg
)

// Wrappers for aurora background colors.
const (
	BlackBg = (1 + iota) << 16
	RedBg
	GreenBg
	BrownBg
	BlueBg
	MagentaBg
	CyanBg
	GrayBg
)

// New creates a new Logger based on the arguments. An empty New() will return a Logger with default settings. Optional arguments are called with following format New(colored, showtimestamp). This is effectively the same as making a new Logger and then calling logger.ShowColor(true) and logger.ShowTimestamp(true).
func New(a ...bool) *Logger {
	c := true
	ts := true
	if len(a) != 0 {
		if !a[0] {
			ts = false
		}
		if len(a) != 1 {
			if !a[1] {
				c = false
			}
		}
	}

	l := Logger{}
	l = Logger{
		Normal,
		ts,
		c,
		aurora.NewAurora(c),
		false,
		"",
		Event{&l, true, true, GreenFg, ShortDate | Time12Hour | TimeZone, Prefix, "DEBUG:"},
		Event{&l, true, true, GrayFg, ShortDate | Time12Hour | TimeZone, Prefix, "INFO:"},
		Event{&l, true, true, YellowFg, ShortDate | Time12Hour | TimeZone, Prefix, "NOTICE:"},
		Event{&l, true, true, RedFg, ShortDate | Time12Hour | TimeZone, Prefix, "ERROR:"},
	}

	return &l
}

// LogLevel returns the current log level.
func (l *Logger) LogLevel() LogLevel {
	return l.logLevel
}

// SetLogLevel sets the logLevel to the given LogLevel.
func (l *Logger) SetLogLevel(lv LogLevel) {
	l.logLevel = lv
}

// ShowTimestamp sets whether or not to show timestamps for the entire logger.
func (l *Logger) ShowTimestamp(b bool) {
	l.timestamp = b
}

// ShowColor sets whether or not to use colors for the entire logger.
func (l *Logger) ShowColor(b bool) {
	l.colored = b
}

// SaveLog will save the log to a file on disk at the given path.
func (l *Logger) SaveLog(path string) error {
	logFile := "/log.log"
	if !l.toDisk {
		l.toDisk = true
	}

	_, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}

	_, err = os.OpenFile(path+logFile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	l.logPath = path + logFile
	return nil
}

// StopSaveLog will stop logging from happening with the current logger.
func (l *Logger) StopSaveLog() {
	l.toDisk = false
}

// validateTimestamp returns true if the given timestamp format is valid.
func validateTimestamp(timestamp int) bool {
	d := true
	h := true
	if (timestamp & datemask) == datemask {
		d = false
	}
	if (timestamp & hourmask) == hourmask {
		h = false
	}

	return d && h
}
