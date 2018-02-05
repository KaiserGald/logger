// Package logger
// 5 February, 2018
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

	"github.com/logrusorgru/aurora"
)

// An Event represents a message with a given level of importance to be printed to the
// log.
type Event struct {
	*Logger   // a Pointer to the parent Logger
	timestamp bool
	colored   bool
	colors    aurora.Color
	format    int
	cformat   ColorFormat
	prefix    string
}

// ShowTimestamp sets whether or not to show timestamps for this log event.
func (e *Event) ShowTimestamp(b bool) {
	e.timestamp = b
}

// ShowColor sets wether or not to show color for this log event.
func (e *Event) ShowColor(b bool) {
	e.colored = b
}

// SetColors sets the foreground color, background color, and special format of the log event
func (e *Event) SetColors(colors aurora.Color) {
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

// SetColorFormat sets the format for the colored output. Timestamp adds color to the timestamp. Prefix adds color to the Prefix. Message adds color to the Message.
func (e *Event) SetColorFormat(format ColorFormat) error {
	if (format | cformatMask) != cformatMask {
		return errors.New("Invalid color format")
	}
	e.cformat = format
	return nil
}

// Prefix returns the prefix of the log event.
func (e *Event) Prefix() string {
	return e.prefix
}

// Log logs the given message via the appropriate log event to STDERR. It will not
// display any log event that is lower than the given level. Debug will not show when
// the log level is Normal.
func (e *Event) Log(fstring string, a ...interface{}) (string, error) {
	var entry string
	var err error
	switch e.Prefix() {
	case "DEBUG:":
		if e.Logger.LogLevel() == All {
			entry, err = e.printf(fstring, a...)
			if err != nil {
				return entry, err
			}
		}
	case "INFO:":
		if e.Logger.LogLevel() <= Verbose {
			entry, err = e.printf(fstring, a...)
			if err != nil {
				return entry, err
			}
		}
	case "NOTICE:":
		if e.Logger.LogLevel() <= Normal {
			entry, err = e.printf(fstring, a...)
			if err != nil {
				return entry, err
			}
		}
	case "ERROR:":
		if e.Logger.LogLevel() <= ErrorsOnly {
			entry, err = e.printf(fstring, a...)
			if err != nil {
				return entry, err
			}
		}
	}

	return entry, nil
}

// buildMessage constructs a message using the given input and format code.
func (e *Event) buildMessage(message string) (string, error) {
	timestamp, err := e.buildTimestamp()
	if err != nil {
		return "", err
	}

	prefix := e.Prefix()
	if e.colored && e.Logger.colored {
		if (e.cformat & Prefix) == Prefix {
			prefix = fmt.Sprint(aurora.Colorize(prefix, e.colors))
		}

		if (e.cformat & Message) == Message {
			message = fmt.Sprint(aurora.Colorize(message, e.colors))
		}
	}
	var fmessage string
	if e.Logger.timestamp && e.timestamp {
		fmessage = timestamp + " - " + prefix + " " + message
	} else {
		fmessage = prefix + " " + message
	}
	return fmessage + "\t\n", nil
}

func (e *Event) buildTimestamp() (string, error) {
	var datestamp, timestamp, zone string
	var words []string
	if e.Logger.timestamp && e.timestamp {
		if ok := validateTimestamp(e.format); !ok {
			return "", errors.New("Invalid date flags")
		}
		t := time.Now()
		if (e.format & datemask) == ShortDate {
			datestamp = t.Format("1/2/2006")
		} else if (e.format & datemask) == LongDate {
			datestamp = t.Format("2 Jan 2006")
		}
		if datestamp != "" {
			words = append(words, datestamp)
		}

		if (e.format & hourmask) == Time12Hour {
			timestamp = t.Format("3:04:05 PM")
		} else if (e.format & hourmask) == Time24Hour {
			timestamp = t.Format("15:04:05")
		}
		if timestamp != "" {
			words = append(words, timestamp)
		}

		if (e.format & TimeZone) == TimeZone {
			zone = t.Format("MST")
		}

		if zone != "" {
			words = append(words, zone)
		}
	}
	fstamp := strings.Join(words, " ")
	if (e.cformat&Timestamp) == Timestamp &&
		e.colored && e.Logger.colored {
		fstamp = fmt.Sprint(aurora.Colorize(fstamp, e.colors))
	}
	return fstamp + "\t", nil
}

// prints a message to the stderr
func (e *Event) printf(fstring string, a ...interface{}) (string, error) {
	var spacing int
	message, err := e.buildMessage(fstring)
	if err != nil {
		return "", err
	}

	spacing = e.setSpacing()
	w := tabwriter.NewWriter(os.Stderr, spacing, 0, 0, ' ', 0)
	fmt.Fprint(w, fmt.Sprintf(message, a...))
	w.Flush()
	return fmt.Sprintf(message, a...), nil
}

func (e *Event) setSpacing() int {
	if (e.cformat & Timestamp) == Timestamp {
		return 35
	}
	if (e.cformat & Prefix) == Prefix {
		return 26
	}
	return 0
}
