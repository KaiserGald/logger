// Package logger
// 3 January, 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package logger

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

// LogLevel is an uint8 that corresponds to a logging level
type LogLevel uint8

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
	Debug     Event
	Info      Event
	Notice    Event
	Error     Event
}

// Event contains information about the logging level
type Event struct {
	*Logger
	timestamp bool
	colored   bool
	colors    int
	format    int
	prefix    string
}

// None is no color value
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

// New creates a new Logger based on the arguments. An empty New() will return a Logger with default settings. If
// called with one boolean arg, it will determine whether or not to show timestamps. The second arg is for whether or
// not colorize the output.
func New(a ...bool) *Logger {
	l := Logger{}
	l = Logger{
		Normal,
		true,
		true,
		Event{&l, true, true, GreenFg | None | None, ShortDate | Time12Hour | TimeZone, "DEBUG:"},
		Event{&l, true, true, None | None | None, ShortDate | Time12Hour | TimeZone, "INFO:"},
		Event{&l, true, true, BrownFg | None | None, ShortDate | Time12Hour | TimeZone, "NOTICE:"},
		Event{&l, true, true, RedFg | None | None, ShortDate | Time12Hour | TimeZone, "ERROR:"},
	}

	if len(a) != 0 {
		if !a[0] {
			l.timestamp = false
		}
		if len(a) != 1 {
			if !a[1] {
				l.colored = false
			}
		}
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

// ShowTimestamp sets whether or not to show timestamps for this log event.
func (e *Event) ShowTimestamp(b bool) {
	e.timestamp = b
}

// ShowColor sets wether or not to show color for this log event
func (e *Event) ShowColor(b bool) {
	e.colored = b
}

// SetColors sets the foreground color, background color, and special format of the log event
func (e *Event) SetColors(colors int) {
	e.colors = colors
}

// SetFormat sets the format flags for configuring the timestamp of the event
func (e *Event) SetFormat(format int) error {
	if ok := validateTimestamp(format); !ok {
		return errors.New("Invalid format flag combination")
	}
	e.format = format
	return nil
}

// Prefix returns the prefix of the log event
func (e *Event) Prefix() string {
	return e.prefix
}

// Log logs the given message via the appropriate log event to STDERR
func (e *Event) Log(fstring string, a ...interface{}) (string, error) {
	var entry string
	var err error
	switch e.Prefix() {
	case "DEBUG:":
		if e.Logger.LogLevel() == All {
			entry, err = e.printf(e.format, fstring, a...)
			if err != nil {
				return entry, err
			}
		}
	case "INFO:":
		if e.Logger.LogLevel() <= Verbose {
			entry, err = e.printf(e.format, fstring, a...)
			if err != nil {
				return entry, err
			}
		}
	case "NOTICE:":
		if e.Logger.LogLevel() <= Normal {
			entry, err = e.printf(e.format, fstring, a...)
			if err != nil {
				return entry, err
			}
		}
	case "ERROR:":
		if e.Logger.LogLevel() <= ErrorsOnly {
			entry, err = e.printf(e.format, fstring, a...)
			if err != nil {
				return entry, err
			}
		}
	}

	return entry, nil
}

// buildMessage constructs a message using the given input and format code
func (e *Event) buildMessage(message string, format int) (string, error) {
	timestamp, err := e.buildTimestamp(format)
	if err != nil {
		return "", err
	}

	fmessage := timestamp + " - " + e.Prefix() + " " + message
	return fmessage + "\t\n", nil
}

func (e *Event) buildTimestamp(format int) (string, error) {
	var datestamp, timestamp, zone string
	var words []string
	if e.Logger.timestamp && e.timestamp {
		if ok := validateTimestamp(format); !ok {
			return "", errors.New("Invalid date flags")
		}
		t := time.Now()
		if (format & datemask) == ShortDate {
			datestamp = t.Format("1/2/2006")
		} else if (format & datemask) == LongDate {
			datestamp = t.Format("2 Jan 2006")
		}
		if datestamp != "" {
			words = append(words, datestamp)
		}

		if (format & hourmask) == Time12Hour {
			timestamp = t.Format("3:04:05 PM")
		} else if (format & hourmask) == Time24Hour {
			timestamp = t.Format("15:04:05")
		}
		if timestamp != "" {
			words = append(words, timestamp)
		}

		if (format & TimeZone) == TimeZone {
			zone = t.Format("MST")
		}

		if zone != "" {
			words = append(words, zone)
		}
	}
	fstamp := strings.Join(words, " ")
	return fstamp + "\t", nil
}

// prints a message to the stderr
func (e *Event) printf(format int, fstring string, a ...interface{}) (string, error) {
	message, err := e.buildMessage(fstring, format)
	if err != nil {
		return "", err
	}
	w := tabwriter.NewWriter(os.Stderr, 26, 0, 0, ' ', 0)
	fmt.Fprintf(w, message, a...)
	w.Flush()
	return fmt.Sprintf(message, a...), nil
}
