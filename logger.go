// Package logger
// 3 January, 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package logger

import (
	"github.com/logrusorgru/aurora"
)

// LogLevel is an uint8 that corresponds to a logging level
type LogLevel uint8

// ColorFormat is a uint8 for color flags
type ColorFormat uint8

// constants for LogLevels
const (
	All LogLevel = iota
	Verbose
	Normal
	ErrorsOnly
	Test
)

// constants for timestamp format
const (
	ShortDate = 1 << iota
	LongDate
	Time12Hour
	Time24Hour
	TimeZone
	datemask = ShortDate | LongDate
	hourmask = Time12Hour | Time24Hour
	timemask = TimeZone
)

// Logger struct
type Logger struct {
	logLevel  LogLevel
	timestamp bool
	colored   bool
	au        aurora.Aurora
	Debug     Event
	Info      Event
	Notice    Event
	Error     Event
}

// Color format flags
const (
	Timestamp ColorFormat = 1 << iota
	Prefix
	Message
	cformatMask = Timestamp | Prefix | Message
)

// wrappers for aurora special formats
const (
	Bold aurora.Color = 1 << iota
	Inverse
)

// wrappers for aurora foreground colors
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

// wrappers for aurora background colors
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

// New creates a new Logger based on the arguments. An empty New() will return a Logger with default settings. If
// called with one boolean arg, it will determine whether or not to show timestamps. The second arg is for whether or
// not colorize the output.
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
		Event{&l, true, true, GreenFg, ShortDate | Time12Hour | TimeZone, Prefix, "DEBUG:"},
		Event{&l, true, true, GrayFg, ShortDate | Time12Hour | TimeZone, Prefix, "INFO:"},
		Event{&l, true, true, YellowFg, ShortDate | Time12Hour | TimeZone, Prefix, "NOTICE:"},
		Event{&l, true, true, RedFg, ShortDate | Time12Hour | TimeZone, Prefix, "ERROR:"},
	}

	return &l
}

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

// LogLevel returns the current log level
func (l *Logger) LogLevel() LogLevel {
	return l.logLevel
}

// SetLogLevel sets the logLevel
func (l *Logger) SetLogLevel(lv LogLevel) {
	l.logLevel = lv
}

// ShowTimestamp sets whether or not to show timestamps for the entire logger
func (l *Logger) ShowTimestamp(b bool) {
	l.timestamp = b
}

// ShowColor sets whether or not to use colors for the entire logger
func (l *Logger) ShowColor(b bool) {
	l.colored = b
}
