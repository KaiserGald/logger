// package logger
// 31 January, 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package logger

import (
	"testing"
	"time"

	"github.com/logrusorgru/aurora"
)

func TestEventShowTimestamp(t *testing.T) {
	test := New()
	test.Debug.ShowTimestamp(false)
	if test.Debug.timestamp {
		t.Errorf("Timestamp was not properly set, expected '%v' got '%v'", false, test.Debug.timestamp)
	}

	test.Debug.ShowTimestamp(true)
	if !test.Debug.timestamp {
		t.Errorf("Timestamp was not properly set, expected '%v' got '%v'", false, test.Debug.timestamp)
	}

	test.Info.ShowTimestamp(false)
	if test.Info.timestamp {
		t.Errorf("Timestamp was not properly set, expected '%v' got '%v'", false, test.Info.timestamp)
	}

	test.Info.ShowTimestamp(true)
	if !test.Info.timestamp {
		t.Errorf("Timestamp was not properly set, expected '%v' got '%v'", false, test.Info.timestamp)
	}

	test.Notice.ShowTimestamp(false)
	if test.Notice.timestamp {
		t.Errorf("Timestamp was not properly set, expected '%v' got '%v'", false, test.Notice.timestamp)
	}

	test.Notice.ShowTimestamp(true)
	if !test.Notice.timestamp {
		t.Errorf("Timestamp was not properly set, expected '%v' got '%v'", false, test.Notice.timestamp)
	}

	test.Error.ShowTimestamp(false)
	if test.Error.timestamp {
		t.Errorf("Timestamp was not properly set, expected '%v' got '%v'", false, test.Error.timestamp)
	}

	test.Error.ShowTimestamp(true)
	if !test.Error.timestamp {
		t.Errorf("Timestamp was not properly set, expected '%v' got '%v'", false, test.Error.timestamp)
	}
}

func TestEventShowColor(t *testing.T) {
	test := New()
	test.Debug.ShowColor(false)
	if test.Debug.colored {
		t.Errorf("Color was not properly set, expected '%v' got '%v'", false, test.Debug.colored)
	}

	test.Debug.ShowColor(true)
	if !test.Debug.colored {
		t.Errorf("Color was not properly set, expected '%v' got '%v'", false, test.Debug.colored)
	}

	test.Info.ShowColor(false)
	if test.Info.colored {
		t.Errorf("Color was not properly set, expected '%v' got '%v'", false, test.Info.colored)
	}

	test.Info.ShowColor(true)
	if !test.Info.colored {
		t.Errorf("Color was not properly set, expected '%v' got '%v'", false, test.Info.colored)
	}

	test.Notice.ShowColor(false)
	if test.Notice.colored {
		t.Errorf("Color was not properly set, expected '%v' got '%v'", false, test.Notice.colored)
	}

	test.Notice.ShowColor(true)
	if !test.Notice.colored {
		t.Errorf("Color was not properly set, expected '%v' got '%v'", false, test.Notice.colored)
	}

	test.Error.ShowColor(false)
	if test.Error.colored {
		t.Errorf("Color was not properly set, expected '%v' got '%v'", false, test.Error.colored)
	}

	test.Error.ShowColor(true)
	if !test.Error.colored {
		t.Errorf("Color was not properly set, expected '%v' got '%v'", false, test.Error.colored)
	}
}

func TestEventSetColors(t *testing.T) {
	test := New()
	test.Debug.SetColors(Bold | MagentaFg | GrayBg)
	if test.Debug.colors != (Bold | MagentaFg | GrayBg) {
		t.Errorf("Colors were not changed")
	}
	test.Debug.SetColors(GrayFg | MagentaBg)
	if test.Debug.colors != (GrayFg | MagentaBg) {
		t.Errorf("Colors were not changed")
	}
	test.Debug.SetColors(RedFg)
	if test.Debug.colors != (RedFg) {
		t.Errorf("Colors were not changed")
	}
	test.Debug.SetColors(GreenBg)
	if test.Debug.colors != (GreenBg) {
		t.Errorf("Colors were not changed")
	}
	test.Debug.SetColors(Bold)
	if test.Debug.colors != (Bold) {
		t.Errorf("Colors were not changed")
	}
}

func TestEventSetFormat(t *testing.T) {
	test := New()
	test.Debug.SetFormat(LongDate | Time24Hour | TimeZone)
	expected := LongDate | Time24Hour | TimeZone
	if test.Debug.format != expected {
		t.Errorf("Format flags do not match, expected '%v' got '%v'", expected, test.Debug.format)
	}
	if err := test.Debug.SetFormat(LongDate | ShortDate); err == nil {
		t.Errorf("Invalid flag combination did not trigger error")
	}
}

func TestEventSetColorFormat(t *testing.T) {
	test := New()
	if err := test.Debug.SetColorFormat(Prefix); err != nil {
		t.Errorf("Error setting color format: '%v'", err)
	}

	if err := test.Debug.SetColorFormat(123); err == nil {
		t.Errorf("Bad format data didn't trigger error")
	}

}

func TestEventPrefix(t *testing.T) {
	test := New()
	expected := "DEBUG:"
	if test.Debug.Prefix() != expected {
		t.Errorf("Error getting event prefix, expected '%v' got '%v'", expected, test.Debug.Prefix())
	}

	expected = "INFO:"
	if test.Info.Prefix() != expected {
		t.Errorf("Error getting event prefix, expected '%v' got '%v'", expected, test.Info.Prefix())
	}

	expected = "NOTICE:"
	if test.Notice.Prefix() != expected {
		t.Errorf("Error getting event prefix, expected '%v' got '%v'", expected, test.Notice.Prefix())
	}

	expected = "ERROR:"
	if test.Error.Prefix() != expected {
		t.Errorf("Error getting event prefix, expected '%v' got '%v'", expected, test.Error.Prefix())
	}
}

func TestEventLog(t *testing.T) {
	greenfg := esc + aurora.GreenFg.Nos() + "m"
	grayfg := esc + aurora.GrayFg.Nos() + "m"
	yellowfg := esc + aurora.BrownFg.Nos() + "m"
	redfg := esc + aurora.RedFg.Nos() + "m"
	test := New()
	message := "Test message"
	test.SetLogLevel(All)
	res, err := test.Debug.Log(message)
	if err != nil {
		t.Errorf("Error building string: %v", err)
	}
	tn := time.Now()
	tf := tn.Format("1/2/2006 3:04:05 PM MST")
	expected := tf + " - " + greenfg + test.Debug.Prefix() + clear + " " + message
	res = trimSpaces(res)
	expected = trimSpaces(expected)
	if res != expected {
		t.Errorf("Strings do not match, expected '%v' got '%v'", expected, res)
	}
	test.SetLogLevel(Normal)
	res, err = test.Debug.Log(message)
	if err != nil {
		t.Errorf("Error building string: %v", err)
	}
	expected = ""
	if res != expected {
		t.Errorf("Strings do not match, expected '%v' got '%v'", expected, res)
	}

	test.SetLogLevel(Verbose)
	test.Info.SetColorFormat(Timestamp)
	test.Info.SetFormat(Time24Hour | TimeZone)
	res, err = test.Info.Log(message)
	if err != nil {
		t.Errorf("Error building string: %v", err)
	}
	tn = time.Now()
	tf = tn.Format("15:04:05 MST")
	expected = grayfg + tf + clear + " - " + test.Info.Prefix() + " " + message
	res = trimSpaces(res)
	expected = trimSpaces(expected)
	if res != expected {
		t.Errorf("Strings do not match, expected '%v' got '%v'", expected, res)
	}

	test.SetLogLevel(Normal)
	test.Notice.SetFormat(LongDate | TimeZone)
	test.Notice.SetColorFormat(Timestamp | Prefix)
	res, err = test.Notice.Log(message)
	if err != nil {
		t.Errorf("Error building string: %v", err)
	}
	tn = time.Now()
	tf = tn.Format("2 Jan 2006 MST")
	expected = yellowfg + tf + clear + " - " + yellowfg + test.Notice.Prefix() + clear + " " + message
	res = trimSpaces(res)
	expected = trimSpaces(expected)
	if res != expected {
		t.Errorf("Strings do not match, expected '%v' got '%v'", expected, res)
	}

	test.SetLogLevel(ErrorsOnly)
	test.Error.SetFormat(LongDate | Time12Hour | TimeZone)
	test.Error.SetColorFormat(Message)
	res, err = test.Error.Log(message)
	if err != nil {
		t.Errorf("Error building string: %v", err)
	}
	tn = time.Now()
	tf = tn.Format("2 Jan 2006 3:04:05 PM MST")
	expected = tf + " - " + test.Error.Prefix() + " " + redfg + message + clear
	res = trimSpaces(res)
	expected = trimSpaces(expected)
	if res != expected {
		t.Errorf("Strings do not match, expected '%v' got '%v'", expected, res)
	}

	test.SetLogLevel(All)
	test.Debug.SetFormat(ShortDate | Time12Hour | TimeZone)
	test.Debug.SetColorFormat(Prefix | Message)
	res, err = test.Debug.Log(message)
	if err != nil {
		t.Errorf("Error building string: %v", err)
	}
	tn = time.Now()
	tf = tn.Format("1/2/2006 3:04:05 PM MST")
	expected = tf + " - " + greenfg + test.Debug.Prefix() + clear + " " + greenfg + message + clear
	res = trimSpaces(res)
	expected = trimSpaces(expected)
	if res != expected {
		t.Errorf("Strings do not match, expected '%v' got '%v'", expected, res)
	}

	test.Error.SetColorFormat(Timestamp | Prefix | Message)
	test.Error.SetFormat(ShortDate | Time12Hour | TimeZone)
	res, err = test.Error.Log(message)
	if err != nil {
		t.Errorf("Error building string: %v", err)
	}
	tn = time.Now()
	tf = tn.Format("1/2/2006 3:04:05 PM MST")
	expected = redfg + tf + clear + " - " + redfg + test.Error.Prefix() + clear + " " + redfg + message + clear
	res = trimSpaces(res)
	expected = trimSpaces(expected)
	if res != expected {
		t.Errorf("Strings do not match, expected '%v' got '%v'", expected, res)
	}

	test.Info.SetColorFormat(Timestamp | Message)
	test.Info.SetFormat(ShortDate | Time12Hour | TimeZone)
	res, err = test.Info.Log(message)
	if err != nil {
		t.Errorf("Error building string: %v", err)
	}
	tn = time.Now()
	tf = tn.Format("1/2/2006 3:04:05 PM MST")
	expected = grayfg + tf + clear + " - " + test.Info.Prefix() + " " + grayfg + message + clear
	res = trimSpaces(res)
	expected = trimSpaces(expected)
	if res != expected {
		t.Errorf("Strings do not match, expected '%v' got '%v'", expected, res)
	}

	test.ShowColor(false)
	res, err = test.Debug.Log(message)
	if err != nil {
		t.Errorf("Error building string: %v", err)
	}
	tn = time.Now()
	tf = tn.Format("1/2/2006 3:04:05 PM MST")
	expected = tf + " - " + test.Debug.Prefix() + " " + message
	res = trimSpaces(res)
	expected = trimSpaces(expected)
	if res != expected {
		t.Errorf("Strings do not match, expected '%v' got '%v'", expected, res)
	}

	test.Debug.format = ShortDate | LongDate
	test.Info.format = ShortDate | LongDate
	test.Notice.format = ShortDate | LongDate
	test.Error.format = ShortDate | LongDate

	if _, err = test.Debug.Log(message); err == nil {
		t.Errorf("Bad format flags did not trigger error")
	}
	if _, err = test.Info.Log(message); err == nil {
		t.Errorf("Bad format flags did not trigger error")
	}
	if _, err = test.Notice.Log(message); err == nil {
		t.Errorf("Bad format flags did not trigger error")
	}
	if _, err = test.Error.Log(message); err == nil {
		t.Errorf("Bad format flags did not trigger error")
	}
}

func TestEventBuildMessage(t *testing.T) {
	greenfg := esc + aurora.GreenFg.Nos() + "m"
	test := New()
	now := time.Now()
	timeformat := now.Format("1/2/2006 3:04:05 PM MST")
	expected := timeformat + " - " + greenfg + "DEBUG:" + clear + " Test event"
	actual, err := test.Debug.buildMessage("Test event")
	if err != nil {
		t.Errorf("Error building message: %v", err)
	}
	actual = trimSpaces(actual)
	expected = trimSpaces(expected)
	if actual != expected {
		t.Errorf("Messages don't match, expected '%v' got '%v'", expected, actual)
	}

	test.Debug.format = ShortDate | LongDate
	_, err = test.Debug.buildMessage("Test event")
	if err == nil {
		t.Errorf("Bad input did not trigger error")
	}
}

func TestEventValidateTimestamp(t *testing.T) {
	expected := false
	actual := validateTimestamp(ShortDate | LongDate)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}

	expected = false
	actual = validateTimestamp(Time12Hour | Time24Hour)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}

	expected = false
	actual = validateTimestamp(ShortDate | LongDate | Time12Hour)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}

	expected = false
	actual = validateTimestamp(ShortDate | Time12Hour | Time24Hour)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}

	expected = false
	actual = validateTimestamp(ShortDate | LongDate | Time12Hour | Time24Hour)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}

	expected = false
	actual = validateTimestamp(ShortDate | LongDate | Time24Hour | TimeZone)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}

	expected = false
	actual = validateTimestamp(ShortDate | LongDate | Time12Hour | Time24Hour | TimeZone)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}

	expected = true
	actual = validateTimestamp(ShortDate | Time12Hour | TimeZone)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}

	expected = true
	actual = validateTimestamp(LongDate | Time24Hour | TimeZone)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}

	expected = true
	actual = validateTimestamp(LongDate)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}

	expected = true
	actual = validateTimestamp(Time12Hour)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}

	expected = true
	actual = validateTimestamp(LongDate | TimeZone)
	if actual != expected {
		t.Errorf("Mismatched results, expected '%v' got '%v'", expected, actual)
	}
}

func TestEventBuildTimestamp(t *testing.T) {
	test := New()
	now := time.Now()
	expectedf := now.Format("1/2/2006")
	test.Debug.SetFormat(ShortDate)
	actualf, err := test.Debug.buildTimestamp()
	if err != nil {
		t.Errorf("Error building timestamp: %v", err)
	}
	actualf = trimSpaces(actualf)

	if actualf != expectedf {
		t.Errorf("Date doesn't match, expected '%v' got '%v'", expectedf, actualf)
	}

	now = time.Now()
	expectedf = now.Format("2 Jan 2006")
	test.Debug.SetFormat(LongDate)
	actualf, err = test.Debug.buildTimestamp()
	if err != nil {
		t.Errorf("Error building timestamp: %v", err)
	}
	actualf = trimSpaces(actualf)
	if actualf != expectedf {
		t.Errorf("Date doesn't match, expected '%v' got '%v'", expectedf, actualf)
	}

	now = time.Now()
	expectedf = now.Format("3:04:05 PM")
	test.Debug.SetFormat(Time12Hour)
	actualf, err = test.Debug.buildTimestamp()
	if err != nil {
		t.Errorf("Error building timestamp: %v", err)
	}
	actualf = trimSpaces(actualf)
	if actualf != expectedf {
		t.Errorf("Date doesn't match, expected '%v' got '%v'", expectedf, actualf)
	}

	now = time.Now()
	expectedf = now.Format("15:04:05")
	test.Debug.SetFormat(Time24Hour)
	actualf, err = test.Debug.buildTimestamp()
	if err != nil {
		t.Errorf("Error building timestamp: %v", err)
	}
	actualf = trimSpaces(actualf)
	if actualf != expectedf {
		t.Errorf("Date doesn't match, expected '%v' got '%v'", expectedf, actualf)
	}

	now = time.Now()
	expectedf = now.Format("3:04:05 PM MST")
	test.Debug.SetFormat(Time12Hour | TimeZone)
	actualf, err = test.Debug.buildTimestamp()
	if err != nil {
		t.Errorf("Error building timestamp: %v", err)
	}
	actualf = trimSpaces(actualf)
	if actualf != expectedf {
		t.Errorf("Date doesn't match, expected '%v' got '%v'", expectedf, actualf)
	}

	now = time.Now()
	test.Debug.format = (ShortDate | LongDate)
	_, err = test.Debug.buildTimestamp()
	if err == nil {
		t.Errorf("Bad input did not trigger error")
	}

}
