// package logger
// 31 January, 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/logrusorgru/aurora"
)

const (
	esc   = "\033["
	clear = esc + "0m"
)

func TestNew(t *testing.T) {
	defexpected := Logger{}
	defexpected = Logger{
		Normal,
		true,
		true,
		aurora.NewAurora(true),
		false,
		"",
		Event{&defexpected, true, true, GreenFg, ShortDate | Time12Hour | TimeZone, Prefix, "DEBUG:"},
		Event{&defexpected, true, true, GrayFg, ShortDate | Time12Hour | TimeZone, Prefix, "INFO:"},
		Event{&defexpected, true, true, YellowFg, ShortDate | Time12Hour | TimeZone, Prefix, "NOTICE:"},
		Event{&defexpected, true, true, RedFg, ShortDate | Time12Hour | TimeZone, Prefix, "ERROR:"},
	}

	defactual := New()
	if !reflect.DeepEqual(defactual, &defexpected) {
		t.Errorf("Default logger not created correctly")
	}

	ntsexpected := Logger{}
	ntsexpected = Logger{
		Normal,
		false,
		true,
		aurora.NewAurora(true),
		false,
		"",
		Event{&ntsexpected, true, true, GreenFg, ShortDate | Time12Hour | TimeZone, Prefix, "DEBUG:"},
		Event{&ntsexpected, true, true, GrayFg, ShortDate | Time12Hour | TimeZone, Prefix, "INFO:"},
		Event{&ntsexpected, true, true, YellowFg, ShortDate | Time12Hour | TimeZone, Prefix, "NOTICE:"},
		Event{&ntsexpected, true, true, RedFg, ShortDate | Time12Hour | TimeZone, Prefix, "ERROR:"},
	}

	ntsactual := New(false)

	if !reflect.DeepEqual(ntsactual, &ntsexpected) {
		t.Errorf("No timestamp logger not created correctly")
	}

	ncexpected := Logger{}
	ncexpected = Logger{
		Normal,
		true,
		false,
		aurora.NewAurora(false),
		false,
		"",
		Event{&ncexpected, true, true, GreenFg, ShortDate | Time12Hour | TimeZone, Prefix, "DEBUG:"},
		Event{&ncexpected, true, true, GrayFg, ShortDate | Time12Hour | TimeZone, Prefix, "INFO:"},
		Event{&ncexpected, true, true, YellowFg, ShortDate | Time12Hour | TimeZone, Prefix, "NOTICE:"},
		Event{&ncexpected, true, true, RedFg, ShortDate | Time12Hour | TimeZone, Prefix, "ERROR:"},
	}

	ncactual := New(true, false)

	if !reflect.DeepEqual(ncactual, &ncexpected) {
		t.Errorf("No color logger not created correctly")
	}

	falseexpected := Logger{}
	falseexpected = Logger{
		Normal,
		false,
		false,
		aurora.NewAurora(false),
		false,
		"",
		Event{&falseexpected, true, true, GreenFg, ShortDate | Time12Hour | TimeZone, Prefix, "DEBUG:"},
		Event{&falseexpected, true, true, GrayFg, ShortDate | Time12Hour | TimeZone, Prefix, "INFO:"},
		Event{&falseexpected, true, true, YellowFg, ShortDate | Time12Hour | TimeZone, Prefix, "NOTICE:"},
		Event{&falseexpected, true, true, RedFg, ShortDate | Time12Hour | TimeZone, Prefix, "ERROR:"},
	}

	falseactual := New(false, false)
	if !reflect.DeepEqual(falseactual, &falseexpected) {
		t.Errorf("No options logger not created correctly")
	}
}

func TestLoggerLogLevel(t *testing.T) {
	test := New()
	lv := test.LogLevel()
	if lv != Normal {
		t.Errorf("Log level did not match test case, expected '%v' got '%v'", Normal, lv)
	}
}

func TestLoggerSetLogLevel(t *testing.T) {
	test := New()
	expected := All
	test.SetLogLevel(All)
	if test.LogLevel() != expected {
		t.Errorf("Log level was not set, expected '%v' got '%v'", All, test.LogLevel())
	}

	expected = Verbose
	test.SetLogLevel(Verbose)
	if test.LogLevel() != expected {
		t.Errorf("Log level was not set, expected '%v' got '%v'", Verbose, test.LogLevel())
	}

	expected = ErrorsOnly
	test.SetLogLevel(ErrorsOnly)
	if test.LogLevel() != expected {
		t.Errorf("Log level was not set, expected '%v' got '%v'", ErrorsOnly, test.LogLevel())
	}

	expected = Test
	test.SetLogLevel(Test)
	if test.LogLevel() != expected {
		t.Errorf("Log level was not set, expected '%v' got '%v'", Test, test.LogLevel())
	}
}

func TestLoggerShowTimestamp(t *testing.T) {
	test := New()
	test.ShowTimestamp(false)
	if test.timestamp {
		t.Errorf("Timestamp was not properly set, expected '%v' got '%v'", false, test.timestamp)
	}

	test.ShowTimestamp(true)
	if !test.timestamp {
		t.Errorf("Timestamp was not properly set, expected '%v' got '%v'", true, test.timestamp)
	}
}

func TestLoggerShowColor(t *testing.T) {
	test := New()
	test.ShowColor(false)
	if test.colored {
		t.Errorf("Color flag was not properly set, expected '%v' got '%v'", false, test.colored)
	}

	test.ShowColor(true)
	if !test.colored {
		t.Errorf("Color flag was not properly set, expected '%v' got '%v'", true, test.colored)
	}
}

func TestLoggerSaveLog(t *testing.T) {
	message := "Test message"
	test := New()
	if err := test.SaveLog("log"); err != nil {
		t.Errorf("Error creating save log: %v", err)
	}
	if !test.toDisk {
		t.Errorf("toDisk flag was not properly set, expected '%v' got '%v'", true, test.toDisk)
	}
	_, err := test.Error.Log(message)
	if err != nil {
		t.Errorf("Error logging event: %v", err)
	}
	_, err = os.Stat(test.logPath)
	if err != nil {
		t.Errorf("Error opening log file: %v", err)
	}

	b, err := ioutil.ReadFile(test.logPath)
	if err != nil {
		t.Errorf("Error reading log file: %v", err)
	}
	line := string(b)
	line = trimSpaces(line)
	fmt.Println("Line:", line)
	tn := time.Now()
	tf := tn.Format("1/2/2006 3:04:05 PM MST")
	expected := tf + " - " + test.Error.Prefix() + " " + message
	expected = trimSpaces(expected)
	if line != expected {
		t.Errorf("Strings do not match, expected '%v' got '%v'", expected, line)
	}

	os.RemoveAll("log")
}

func TestLoggerStopSaveLog(t *testing.T) {
	test := New()
	test.SaveLog("log")
	test.StopSaveLog()
	if test.toDisk {
		t.Errorf("toDisk flag no set, expected '%v' got '%v'", false, test.toDisk)
	}

	os.RemoveAll("log")
}

func trimSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
