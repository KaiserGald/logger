// Package logger
// 3 January, 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package logger

// LogLevel is an uint8 that corresponds to a logging level
type LogLevel uint8

// constants for LogLevels
const (
	All LogLevel = iota + 1
	Verbose
	Normal
	ErrorsOnly
	Test
)

// Logger struct
type Logger struct {
	logLevel  LogLevel
	timestamp bool
	colored   bool
	Debug     Event
	Info      Event
	Notice    Event
	Error     Event
}

// Event contains information about the logging level
type Event struct {
	timestamp bool
	colored   bool
	colors    int
	prefix    string
}

const None int = 0

// wrappers for aurora special formats
const (
	Bold int = 1 << iota
	Inverse
)

// wrappers for aurora foreground colors
const (
	BlackFg int = (1 + iota) << 16
	RedFg
	GreenFg
	BrownFg
	BlueFg
	MagentaFg
	CyanFg
	GrayFg
)

// wrappers for aurora background colors
const (
	BlackBg int = (1 + iota) << 8
	RedBg
	GreenBg
	BrownBg
	BlueBg
	MagentaBg
	CyanBg
	GrayBg
)

// New creates a new Logger based on the arguments. An empty New() will return a Logger with default settings. If called with one boolean arg, it will determine whether or not to show timestamps. The second arg is for whether or not colorize the output.
func New(a ...bool) *Logger {
	l := Logger{}
	return &l
}

// LogLevel returns the current log level
func (l *Logger) LogLevel() LogLevel {
	return l.logLevel
}

// SetLogLevel sets the logLevel
func (l *Logger) SetLogLevel(lv LogLevel) {
}

// ShowTimestamp sets whether or not to show timestamps for the entire logger
func (l *Logger) ShowTimestamp(b bool) {

}

// ShowColor sets whether or not to use colors for the entire logger
func (l *Logger) ShowColor(b bool) {

}

// ShowTimestamp sets whether or not to show timestamps for this log event.
func (e *Event) ShowTimestamp(b bool) {

}

// ShowColor sets wether or not to show color for this log event
func (e *Event) ShowColor(b bool) {

}

// SetColors sets the foreground color, background color, and special format of the log event
func (e *Event) SetColors() {

}

// Prefix returns the prefix of the log event
func (e *Event) Prefix() string {
	return ""
}

// Log logs the given message via the appropriate log event to STDERR
func (e *Event) Log() {

}
