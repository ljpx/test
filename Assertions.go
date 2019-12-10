package test

import (
	"fmt"
	"reflect"
)

// Assertions defines a number of assertions that can be made about x.
type Assertions struct {
	t T
	x interface{}
}

// IsEqualTo fails the test if y is not equal to the subject, x.
func (a *Assertions) IsEqualTo(y interface{}) {
	a.t.Helper()

	if baseEqualityTest(a.x, y) {
		return
	}

	formattedFailure(a.t, "Expected %v to be equal to %v", a.x, y)
}

// IsNotEqualTo fails the test if y is equal to the subject, x.
func (a *Assertions) IsNotEqualTo(y interface{}) {
	a.t.Helper()

	if !baseEqualityTest(a.x, y) {
		return
	}

	formattedFailure(a.t, "Expected %v to not be equal to %v", a.x, y)
}

// IsNil fails the test if the subject, x, is not nil.
func (a *Assertions) IsNil() {
	a.t.Helper()

	if baseNilTest(a.x) {
		return
	}

	formattedFailure(a.t, "Expected %v to be <nil>", a.x)
}

// IsNotNil fails the test if the subject, x, is nil.
func (a *Assertions) IsNotNil() {
	a.t.Helper()

	if !baseNilTest(a.x) {
		return
	}

	formattedFailure(a.t, "Expected subject to not be <nil>, but was")
}

func baseEqualityTest(x interface{}, y interface{}) bool {
	return x == y
}

func baseNilTest(x interface{}) bool {
	if x == nil {
		return true
	}

	v := reflect.ValueOf(x)
	if (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) && v.IsNil() {
		return true
	}

	return false
}

func formattedFailure(t T, format string, args ...interface{}) {
	t.Helper()

	name := t.Name()
	t.Fatalf("\n\n× %v\n%v\n\n", name, fmt.Sprintf(format, args...))
}
