package test

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestIsEqualTo(t *testing.T) {
	testCases := []struct {
		x    interface{}
		y    interface{}
		pass bool
	}{
		{x: 4, y: 3, pass: false},
		{x: 4, y: 4, pass: true},
		{x: "Hello", y: 4, pass: false},
		{x: "Hello", y: "Hello", pass: true},
		{x: "Hellp", y: "Hello", pass: false},
		{x: nil, y: 4, pass: false},
		{x: nil, y: nil, pass: true},
		{x: nil, y: io.Writer(nil), pass: true},
	}

	for _, testCase := range testCases {
		recorder := NewRecorder()
		That(recorder, testCase.x).IsEqualTo(testCase.y)

		if !testCase.pass {
			assertFailed(t, recorder)
			assertFailureMessage(t, recorder, "Expected %v to be equal to %v", testCase.x, testCase.y)
			assertHelperCount(t, recorder, 3)
		} else {
			assertPassed(t, recorder)
			assertHelperCount(t, recorder, 2)
		}
	}
}

func TestIsNotEqualTo(t *testing.T) {
	testCases := []struct {
		x    interface{}
		y    interface{}
		pass bool
	}{
		{x: 4, y: 3, pass: true},
		{x: 4, y: 4, pass: false},
		{x: "Hello", y: 4, pass: true},
		{x: "Hello", y: "Hello", pass: false},
		{x: "Hellp", y: "Hello", pass: true},
		{x: nil, y: 4, pass: true},
		{x: nil, y: nil, pass: false},
		{x: nil, y: io.Writer(nil), pass: false},
	}

	for _, testCase := range testCases {
		recorder := NewRecorder()
		That(recorder, testCase.x).IsNotEqualTo(testCase.y)

		if !testCase.pass {
			assertFailed(t, recorder)
			assertFailureMessage(t, recorder, "Expected %v to not be equal to %v", testCase.x, testCase.y)
			assertHelperCount(t, recorder, 3)
		} else {
			assertPassed(t, recorder)
			assertHelperCount(t, recorder, 2)
		}
	}
}

func TestIsNil(t *testing.T) {
	testCases := []struct {
		x    interface{}
		pass bool
	}{
		{x: 4, pass: false},
		{x: "Hello", pass: false},
		{x: io.Writer(nil), pass: true},
		{x: nil, pass: true},
		{x: (*int)(nil), pass: true},
	}

	for _, testCase := range testCases {
		recorder := NewRecorder()
		That(recorder, testCase.x).IsNil()

		if !testCase.pass {
			assertFailed(t, recorder)
			assertFailureMessage(t, recorder, "Expected %v to be <nil>", testCase.x)
			assertHelperCount(t, recorder, 3)
		} else {
			assertPassed(t, recorder)
			assertHelperCount(t, recorder, 2)
		}
	}
}

func TestIsNotNil(t *testing.T) {
	testCases := []struct {
		x    interface{}
		pass bool
	}{
		{x: 4, pass: true},
		{x: "Hello", pass: true},
		{x: io.Writer(nil), pass: false},
		{x: nil, pass: false},
		{x: (*int)(nil), pass: false},
	}

	for _, testCase := range testCases {
		recorder := NewRecorder()
		That(recorder, testCase.x).IsNotNil()

		if !testCase.pass {
			assertFailed(t, recorder)
			assertFailureMessage(t, recorder, "Expected subject to not be <nil>, but was")
			assertHelperCount(t, recorder, 3)
		} else {
			assertPassed(t, recorder)
			assertHelperCount(t, recorder, 2)
		}
	}
}

func assertFailed(t *testing.T, recorder *Recorder) {
	t.Helper()

	if !recorder.DidFail {
		t.Fatalf("expected the test to fail, but didn't fail")
	}
}

func assertPassed(t *testing.T, recorder *Recorder) {
	t.Helper()

	if recorder.DidFail {
		t.Fatalf("expected the test to pass, but failed")
	}
}

func assertFailureMessage(t *testing.T, recorder *Recorder, format string, args ...interface{}) {
	t.Helper()
	message := fmt.Sprintf(format, args...)

	if recorder.FailMessage == "" {
		t.Fatalf("expected failure message to be like '%v' but was empty", message)
	}

	spl := strings.SplitN(recorder.FailMessage, "× Recorder", 2)
	if len(spl) != 2 {
		t.Fatalf("invalid failure message '%v'", recorder.FailMessage)
	}

	failMessage := strings.TrimSpace(spl[1])
	if failMessage != message {
		t.Fatalf("expected failure message to be like '%v' but was like '%v'", message, failMessage)
	}
}

func assertHelperCount(t *testing.T, recorder *Recorder, count int) {
	t.Helper()

	if recorder.HelperCallCount != count {
		t.Fatalf("expected the helper call count to be %v but was %v", count, recorder.HelperCallCount)
	}
}
