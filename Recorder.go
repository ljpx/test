package test

import "fmt"

// Recorder is a mock implementation of T that is used for unit tests in this
// package.
type Recorder struct {
	HelperCallCount int
	DidFail         bool
	FailMessage     string
}

var _ T = &Recorder{}

// NewRecorder creates a new test recorder.
func NewRecorder() *Recorder {
	return &Recorder{}
}

// Name will always return "Recorder".
func (*Recorder) Name() string {
	return "Recorder"
}

// Helper simply increments `HelperCallCount`.
func (r *Recorder) Helper() {
	r.HelperCallCount++
}

// Fatalf sets `didFail` to true and sets `failMessage` to the failure message.
func (r *Recorder) Fatalf(format string, args ...interface{}) {
	r.DidFail = true
	r.FailMessage = fmt.Sprintf(format, args...)
}
