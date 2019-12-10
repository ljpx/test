package test

import "testing"

func TestRecorderName(t *testing.T) {
	// Arrange.
	sut := NewRecorder()

	// Act.
	name := sut.Name()

	// Assert.
	if name != "Recorder" {
		t.Fatalf("Expected Name() to always return 'Recorder' but return '%v'", name)
	}
}

func TestRecorderHelper(t *testing.T) {
	// Arrange.
	sut := NewRecorder()

	// Precondition.
	if sut.HelperCallCount != 0 {
		t.Fatalf("Initial value for HelperCallCount was not 0")
	}

	// Act.
	sut.Helper()
	sut.Helper()

	// Assert.
	if sut.HelperCallCount != 2 {
		t.Fatalf("Expected two calls of Helper() to produce HelperCallCount == 2 but was actually %v", sut.HelperCallCount)
	}
}

func TestRecorderFatalf(t *testing.T) {
	// Arrange.
	sut := NewRecorder()

	// Precondition.
	if sut.DidFail {
		t.Fatalf("Initial value for DidFail was not false")
	}

	if sut.FailMessage != "" {
		t.Fatalf("Initial value for FailMessage was not an empty string")
	}

	// Act.
	sut.Fatalf("Something %v", "went wrong")

	// Assert.
	if !sut.DidFail {
		t.Fatalf("Expected DidFail to be true, but was false")
	}

	if sut.FailMessage != "Something went wrong" {
		t.Fatalf("Expected FailMessage to be 'Something went wrong' but was '%v'", sut.FailMessage)
	}
}
