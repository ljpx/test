![](icon.svg)

# test

[![GoDoc](https://godoc.org/github.com/ljpx/test?status.svg)](https://godoc.org/github.com/ljpx/test)

Package `test` provides a basic assertion helper for unit tests.

## Usage Example

```go
package math

func TestAdd(t *testing.T) {
    a := 5
    b := 3
    c := Add(a, b)

    test.That(t, c).IsEqualTo(8)
}
```
