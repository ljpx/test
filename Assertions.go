package test

import (
	"fmt"
	"reflect"
	"time"
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

// HasEquivalentSequenceTo fails the test if the subject, x, does not have the
// exact same sequence of values as y.  Only for slices.
func (a *Assertions) HasEquivalentSequenceTo(y interface{}) {
	a.t.Helper()

	xt := reflect.TypeOf(a.x)
	yt := reflect.TypeOf(y)
	xv := reflect.ValueOf(a.x)
	yv := reflect.ValueOf(y)

	if xt.Kind() != reflect.Slice || yt.Kind() != reflect.Slice {
		formattedFailure(a.t, "Expected both subject and comparator to be slices, but subject was a %v and comparator was a %v", xt.Kind(), yt.Kind())
		return
	}

	if xt.Elem() != yt.Elem() {
		formattedFailure(a.t, "Expected subject to have type like %v but was %v", yt, xt)
		return
	}

	if xv.Len() != yv.Len() {
		formattedFailure(a.t, "Expected subject to have length %v but had length %v", yv.Len(), xv.Len())
		return
	}

	if !reflect.DeepEqual(a.x, y) {
		formattedFailure(a.t, "Expected sequence of elements in\n\n%v\n\nto be equal to sequence of elements in\n\n%v", a.x, y)
	}
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
	}
}

// IsGreaterThan fails the test if the subject, x, is not greater than y.
func (a *Assertions) IsGreaterThan(y interface{}) {
	a.t.Helper()

	b, ok := baseGreaterThanTest(a.x, y)
	if !ok {
		formattedFailure(a.t, "Expected two comparable types\nx: %v\ny: %v", typeNameFor(a.x), typeNameFor(y))
		return
	}

	if !b {
		formattedFailure(a.t, "Expected %v to be greater than %v", a.x, y)
	}
}

// IsGreaterThanOrEqualTo fails the test if the subject, x, is not greater than
// or equal to y.
func (a *Assertions) IsGreaterThanOrEqualTo(y interface{}) {
	a.t.Helper()

	b, ok := baseGreaterThanOrEqualToTest(a.x, y)
	if !ok {
		formattedFailure(a.t, "Expected two comparable types\nx: %v\ny: %v", typeNameFor(a.x), typeNameFor(y))
		return
	}

	if !b {
		formattedFailure(a.t, "Expected %v to be greater than or equal to %v", a.x, y)
	}
}

// IsLessThan fails the test if the subject, x, is not less than y.
func (a *Assertions) IsLessThan(y interface{}) {
	a.t.Helper()

	b, ok := baseLessThanTest(a.x, y)
	if !ok {
		formattedFailure(a.t, "Expected two comparable types\nx: %v\ny: %v", typeNameFor(a.x), typeNameFor(y))
		return
	}

	if !b {
		formattedFailure(a.t, "Expected %v to be less than %v", a.x, y)
	}
}

// IsLessThanOrEqualTo fails the test if the subject, x, is not less than
// or equal to y.
func (a *Assertions) IsLessThanOrEqualTo(y interface{}) {
	a.t.Helper()

	b, ok := baseLessThanOrEqualToTest(a.x, y)
	if !ok {
		formattedFailure(a.t, "Expected two comparable types\nx: %v\ny: %v", typeNameFor(a.x), typeNameFor(y))
		return
	}

	if !b {
		formattedFailure(a.t, "Expected %v to be less than or equal to %v", a.x, y)
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

func baseGreaterThanTest(x interface{}, y interface{}) (bool, bool) {
	xi, ok1 := baseIntegerValue(x)
	yi, ok2 := baseIntegerValue(y)
	if ok1 && ok2 {
		return xi > yi, true
	}

	xf, ok1 := baseFloatingValue(x)
	yf, ok2 := baseFloatingValue(y)
	if ok1 && ok2 {
		return xf > yf, true
	}

	xt, ok1 := baseTimeValue(x)
	yt, ok2 := baseTimeValue(y)
	if ok1 && ok2 {
		return xt.Sub(yt) > time.Nanosecond*0, true
	}

	return false, false
}

func baseGreaterThanOrEqualToTest(x interface{}, y interface{}) (bool, bool) {
	xi, ok1 := baseIntegerValue(x)
	yi, ok2 := baseIntegerValue(y)
	if ok1 && ok2 {
		return xi >= yi, true
	}

	xf, ok1 := baseFloatingValue(x)
	yf, ok2 := baseFloatingValue(y)
	if ok1 && ok2 {
		return xf >= yf, true
	}

	xt, ok1 := baseTimeValue(x)
	yt, ok2 := baseTimeValue(y)
	if ok1 && ok2 {
		return xt.Sub(yt) >= time.Nanosecond*0, true
	}

	return false, false
}

func baseLessThanTest(x interface{}, y interface{}) (bool, bool) {
	xi, ok1 := baseIntegerValue(x)
	yi, ok2 := baseIntegerValue(y)
	if ok1 && ok2 {
		return xi < yi, true
	}

	xf, ok1 := baseFloatingValue(x)
	yf, ok2 := baseFloatingValue(y)
	if ok1 && ok2 {
		return xf < yf, true
	}

	xt, ok1 := baseTimeValue(x)
	yt, ok2 := baseTimeValue(y)
	if ok1 && ok2 {
		return xt.Sub(yt) < time.Nanosecond*0, true
	}

	return false, false
}

func baseLessThanOrEqualToTest(x interface{}, y interface{}) (bool, bool) {
	xi, ok1 := baseIntegerValue(x)
	yi, ok2 := baseIntegerValue(y)
	if ok1 && ok2 {
		return xi <= yi, true
	}

	xf, ok1 := baseFloatingValue(x)
	yf, ok2 := baseFloatingValue(y)
	if ok1 && ok2 {
		return xf <= yf, true
	}

	xt, ok1 := baseTimeValue(x)
	yt, ok2 := baseTimeValue(y)
	if ok1 && ok2 {
		return xt.Sub(yt) <= time.Nanosecond*0, true
	}

	return false, false
}

func baseIntegerValue(x interface{}) (int64, bool) {
	switch n := x.(type) {
	case byte:
		return int64(n), true
	case int8:
		return int64(n), true
	case uint16:
		return int64(n), true
	case int16:
		return int64(n), true
	case uint:
		return int64(n), true
	case int:
		return int64(n), true
	case uint32:
		return int64(n), true
	case int32:
		return int64(n), true
	case int64:
		return int64(n), true
	case uint64:
		return int64(n), true
	case uintptr:
		return int64(n), true
	}

	return 0, false
}

func baseFloatingValue(x interface{}) (float64, bool) {
	switch n := x.(type) {
	case float32:
		return float64(n), true
	case float64:
		return float64(n), true
	}

	return 0, false
}

func baseTimeValue(x interface{}) (time.Time, bool) {
	v, ok := x.(time.Time)
	return v, ok
}

func formattedFailure(t T, format string, args ...interface{}) {
	t.Helper()

	name := t.Name()
	t.Fatalf("\n\n× %v\n%v\n\n", name, fmt.Sprintf(format, args...))
}

func typeNameFor(x interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(x))
}
