// package logger
// 31 January, 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package logger

import "testing"

func TestNew(t *testing.T) {
	defexpected := Logger{
		Normal,
		true,
		true,
		Event{true, true, GreenFg | None | None, "DEBUG: "},
		Event{true, true, None | None | None, "INFO: "},
		Event{true, true, BrownFg | None | None, "NOTICE: "},
		Event{true, true, RedFg | None | None, "ERROR: "},
	}

	defactual := New()

	if defactual != &defexpected {
		t.Errorf("Default logger not created correctly")
	}

	ntsexpected := Logger{
		Normal,
		false,
		true,
		Event{true, true, GreenFg | None | None, "DEBUG: "},
		Event{true, true, None | None | None, "INFO: "},
		Event{true, true, BrownFg | None | None, "NOTICE: "},
		Event{true, true, RedFg | None | None, "ERROR: "},
	}

	ntsactual := New(false)

	if ntsactual != &ntsexpected {
		t.Errorf("No timestamp logger not created correctly")
	}

	ncexpected := Logger{
		Normal,
		true,
		false,
		Event{true, true, GreenFg | None | None, "DEBUG: "},
		Event{true, true, None | None | None, "INFO: "},
		Event{true, true, BrownFg | None | None, "NOTICE: "},
		Event{true, true, RedFg | None | None, "ERROR: "},
	}

	ncactual := New(true, false)

	if ncactual != &ncexpected {
		t.Errorf("No color logger not created correctly")
	}

	falseexpected := Logger{
		Normal,
		false,
		false,
		Event{true, true, GreenFg | None | None, "DEBUG: "},
		Event{true, true, None | None | None, "INFO: "},
		Event{true, true, BrownFg | None | None, "NOTICE: "},
		Event{true, true, RedFg | None | None, "ERROR: "},
	}

	falseactual := New(false, false)
	if falseactual != &falseexpected {
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

}
