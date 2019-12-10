package test

// T defines the methods provided by *testing.T that this package uses.  It
// allows for a mock *testing.T to be used in unit tests for this package.
type T interface {
	Name() string
	Helper()
	Fatalf(format string, args ...interface{})
}
