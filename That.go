package test

// That returns a new *Assertions using the provided *testing.T and subject, x.
func That(t T, x interface{}) *Assertions {
	t.Helper()

	return &Assertions{
		t: t,
		x: x,
	}
}
