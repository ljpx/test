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

	formattedFailure(a.t, "Expected %v to be equal to %v\nx: %v\ny: %v", a.x, y, typeNameFor(a.x), typeNameFor(y))
}

// IsNotEqualTo fails the test if y is equal to the subject, x.
func (a *Assertions) IsNotEqualTo(y interface{}) {
	a.t.Helper()

	if !baseEqualityTest(a.x, y) {
		return
	}

	formattedFailure(a.t, "Expected %v to not be equal to %v\nx: %v\ny: %v", a.x, y, typeNameFor(a.x), typeNameFor(y))
}

// IsNil fails the test if the subject, x, is not nil.
func (a *Assertions) IsNil() {
	a.t.Helper()

	if baseNilTest(a.x) {
		return
	}

	formattedFailure(a.t, "Expected %v to be <nil>\nx: %v", a.x, typeNameFor(a.x))
}

// IsNotNil fails the test if the subject, x, is nil.
func (a *Assertions) IsNotNil() {
	a.t.Helper()

	if !baseNilTest(a.x) {
		return
	}

	formattedFailure(a.t, "Expected subject to not be <nil>, but was\nx: %v", typeNameFor(a.x))
}

// IsTrue fails the test if the subject, x, is not a boolean, or is false.
func (a *Assertions) IsTrue() {
	a.t.Helper()

	b, ok := baseBooleanTest(a.x)
	if !ok {
		formattedFailure(a.t, "Expected <true>, but was not a boolean\nx: %v", typeNameFor(a.x))
		return
	}

	if !b {
		formattedFailure(a.t, "Expected <true>, but was <false>")
		return
	}
}

// IsFalse fails the test if the subject, x, is not a boolean, or is true.
func (a *Assertions) IsFalse() {
	a.t.Helper()

	b, ok := baseBooleanTest(a.x)
	if !ok {
		formattedFailure(a.t, "Expected <false>, but was not a boolean\nx: %v", typeNameFor(a.x))
		return
	}

	if b {
		formattedFailure(a.t, "Expected <false>, but was <true>")
		return
	}
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

func baseBooleanTest(x interface{}) (bool, bool) {
	b, ok := x.(bool)
	if !ok {
		return false, false
	}

	return b, true
}

func formattedFailure(t T, format string, args ...interface{}) {
	t.Helper()

	name := t.Name()
	t.Fatalf("\n\n× %v\n%v\n\n", name, fmt.Sprintf(format, args...))
}

func typeNameFor(x interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(x))
}
