package test

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"time"
)

func TestIsEqualTo(t *testing.T) {
	testCases := []struct {
		x    interface{}
		y    interface{}
		pass bool
	}{
		{x: 4, y: 3, pass: false},
		{x: 4, y: 4, pass: true},
		{x: 4, y: int16(4), pass: false},
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

func TestHasEquivalentSequenceToExpectsSlices(t *testing.T) {
	// Arrange.
	x := 5
	y := struct{}{}

	// Act.
	recorder := NewRecorder()

	That(recorder, x).HasEquivalentSequenceTo(y)

	// Assert.
	assertFailed(t, recorder)
	assertHelperCount(t, recorder, 3)
	assertFailureMessage(t, recorder, "Expected both subject and comparator to be slices, but subject was a int and comparator was a struct")
}

func TestHasEquivalentSequenceToExpectsSameSliceType(t *testing.T) {
	// Arrange.
	x := []int{1, 2, 3}
	y := []byte{1, 2, 3}

	// Act.
	recorder := NewRecorder()

	That(recorder, x).HasEquivalentSequenceTo(y)

	// Assert.
	assertFailed(t, recorder)
	assertHelperCount(t, recorder, 3)
	assertFailureMessage(t, recorder, "Expected subject to have type like []uint8 but was []int")
}

func TestHasEquivalentSequenceToExpectsEqualLength(t *testing.T) {
	// Arrange.
	x := []byte{1, 2, 3}
	y := []byte{1, 2, 3, 4}

	// Act.
	recorder := NewRecorder()

	That(recorder, x).HasEquivalentSequenceTo(y)

	// Assert.
	assertFailed(t, recorder)
	assertHelperCount(t, recorder, 3)
	assertFailureMessage(t, recorder, "Expected subject to have length 4 but had length 3")
}

func TestHasEquivalentSequenceToExpectsSameSequence(t *testing.T) {
	// Arrange.
	x := []byte{1, 2, 3}
	y := []byte{1, 2, 4}

	// Act.
	recorder := NewRecorder()

	That(recorder, x).HasEquivalentSequenceTo(y)

	// Assert.
	assertFailed(t, recorder)
	assertHelperCount(t, recorder, 3)
	assertFailureMessage(t, recorder, "Expected sequence of elements in\n\n[1 2 3]\n\nto be equal to sequence of elements in\n\n[1 2 4]")
}

func TestHasEquivalentSequenceToSuccess(t *testing.T) {
	// Arrange.
	x := []byte{1, 2, 3}
	y := []byte{1, 2, 3}

	// Act.
	recorder := NewRecorder()

	That(recorder, x).HasEquivalentSequenceTo(y)

	// Assert.
	assertPassed(t, recorder)
	assertHelperCount(t, recorder, 2)
}

func TestIsTrueNonBoolean(t *testing.T) {
	testCases := []interface{}{
		4,
		"Hello",
		nil,
		io.Writer(nil),
		(*io.Writer)(nil),
	}

	for _, testCase := range testCases {
		recorder := NewRecorder()
		That(recorder, testCase).IsTrue()

		assertFailed(t, recorder)
		assertFailureMessage(t, recorder, "Expected <true>, but was not a boolean")
		assertHelperCount(t, recorder, 3)
	}
}

func TestIsTrue(t *testing.T) {
	testCases := []struct {
		x    interface{}
		pass bool
	}{
		{x: 4 == 4, pass: true},
		{x: "Hello" == "Hello", pass: true},
		{x: 4 == 5, pass: false},
		{x: true, pass: true},
	}

	for _, testCase := range testCases {
		recorder := NewRecorder()
		That(recorder, testCase.x).IsTrue()

		if !testCase.pass {
			assertFailed(t, recorder)
			assertFailureMessage(t, recorder, "Expected <true>, but was <false>")
			assertHelperCount(t, recorder, 3)
		} else {
			assertPassed(t, recorder)
			assertHelperCount(t, recorder, 2)
		}
	}
}

func TestIsFalseNonBoolean(t *testing.T) {
	testCases := []interface{}{
		4,
		"Hello",
		nil,
		io.Writer(nil),
		(*io.Writer)(nil),
	}

	for _, testCase := range testCases {
		recorder := NewRecorder()
		That(recorder, testCase).IsFalse()

		assertFailed(t, recorder)
		assertFailureMessage(t, recorder, "Expected <false>, but was not a boolean")
		assertHelperCount(t, recorder, 3)
	}
}

func TestIsFalse(t *testing.T) {
	testCases := []struct {
		x    interface{}
		pass bool
	}{
		{x: 4 == 4, pass: false},
		{x: "Hello" == "Hello", pass: false},
		{x: 4 == 5, pass: true},
		{x: true, pass: false},
	}

	for _, testCase := range testCases {
		recorder := NewRecorder()
		That(recorder, testCase.x).IsFalse()

		if !testCase.pass {
			assertFailed(t, recorder)
			assertFailureMessage(t, recorder, "Expected <false>, but was <true>")
			assertHelperCount(t, recorder, 3)
		} else {
			assertPassed(t, recorder)
			assertHelperCount(t, recorder, 2)
		}
	}
}

func TestOrderedComparisons(t *testing.T) {
	fixedTime := time.Now()

	testCases := []struct {
		x   interface{}
		y   interface{}
		gt  bool
		gte bool
		lt  bool
		lte bool
	}{
		{x: 5, y: 5, gt: false, gte: true, lt: false, lte: true},
		{x: 4, y: 5, gt: false, gte: false, lt: true, lte: true},
		{x: 6, y: 5, gt: true, gte: true, lt: false, lte: false},
		{x: 5.5, y: 5.5, gt: false, gte: true, lt: false, lte: true},
		{x: 4.5, y: 5.5, gt: false, gte: false, lt: true, lte: true},
		{x: 6.5, y: 5.5, gt: true, gte: true, lt: false, lte: false},
		{x: uint8(5), y: int64(5), gt: false, gte: true, lt: false, lte: true},
		{x: int16(4), y: byte(5), gt: false, gte: false, lt: true, lte: true},
		{x: rune(6), y: uintptr(5), gt: true, gte: true, lt: false, lte: false},
		{x: float64(5.5), y: float32(5.5), gt: false, gte: true, lt: false, lte: true},
		{x: 4.5, y: float64(5.5), gt: false, gte: false, lt: true, lte: true},
		{x: float32(6.5), y: 5.5, gt: true, gte: true, lt: false, lte: false},
		{x: time.Now().Add(time.Second), y: time.Now(), gt: true, gte: true, lt: false, lte: false},
		{x: time.Now().Add(time.Second), y: time.Now().Add(time.Minute), gt: false, gte: false, lt: true, lte: true},
		{x: fixedTime, y: fixedTime, gt: false, gte: true, lt: false, lte: true},
	}

	for _, testCase := range testCases {
		rgt := NewRecorder()
		rgte := NewRecorder()
		rlt := NewRecorder()
		rlte := NewRecorder()

		That(rgt, testCase.x).IsGreaterThan(testCase.y)
		That(rgte, testCase.x).IsGreaterThanOrEqualTo(testCase.y)
		That(rlt, testCase.x).IsLessThan(testCase.y)
		That(rlte, testCase.x).IsLessThanOrEqualTo(testCase.y)

		if testCase.gt {
			assertPassed(t, rgt)
		} else {
			assertFailed(t, rgt)
		}

		if testCase.gte {
			assertPassed(t, rgte)
		} else {
			assertFailed(t, rgte)
		}

		if testCase.lt {
			assertPassed(t, rlt)
		} else {
			assertFailed(t, rlt)
		}

		if testCase.lte {
			assertPassed(t, rlte)
		} else {
			assertFailed(t, rlte)
		}
	}
}

func TestNonOrderedTypes(t *testing.T) {
	// Arrange.
	recorder := NewRecorder()

	// Act.
	That(recorder, 5.5).IsGreaterThan(5)

	// Assert.
	assertFailed(t, recorder)
	assertFailureMessage(t, recorder, "Expected two comparable types\nx: float64\ny: int")
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
	if !strings.Contains(failMessage, message) {
		t.Fatalf("expected failure message to be like '%v' but was like '%v'", message, failMessage)
	}
}

func assertHelperCount(t *testing.T, recorder *Recorder, count int) {
	t.Helper()

	if recorder.HelperCallCount != count {
		t.Fatalf("expected the helper call count to be %v but was %v", count, recorder.HelperCallCount)
	}
}
