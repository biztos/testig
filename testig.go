// testig.go

// Package testig contains helper functions useful in testing Go programs.
package testig

import (
	"fmt"
	"strings"

	"github.com/stretchr/testify/assert"
)

// TT defines an interface implemented by both the TestTester and testing.T
// (and for that matter testing.B).
//
// In order to be testable within this package, helper functions should accept
// a TT argument instead of a *testing.T:
//
//   func TestHelper(t TT) {
//    ...
//   }
//
// This awful hack is necessary because the Go authors explicitly made it
// impossible to implement the testing.TB interface.
type TT interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
	// ...just NOT the annoying private() method
}

// TestTester implements the TT interface in a way that helps us test
// our test functions.  It should be used to run a single test function.
// NOTE: it does not (yet) actually stop execution!
type TestTester struct {
	Logs    []string
	Stopped bool
	failed  bool
	skipped bool
}

// NewTestTester returns a new TestTester ready for testing tests.
func NewTestTester() *TestTester {
	return &TestTester{
		Logs: []string{},
	}
}

func (tt *TestTester) private() {
	// boo!
}

// Error, as in testing.T, is equivalent to Log followed by Fail.
func (tt *TestTester) Error(args ...interface{}) {
	tt.Log(args...)
	tt.Fail()

}

// Errorf, as in testing.T, is equivalent to Logf followed by Fail.
func (tt *TestTester) Errorf(format string, args ...interface{}) {
	tt.Logf(format, args...)
	tt.Fail()
}

// Fail, as in testing.T, marks the function as having failed but continues
// execution.
func (tt *TestTester) Fail() {
	tt.failed = true
}

// FailNow, as in testing.T, marks the function as having failed and stops
// its execution.
func (tt *TestTester) FailNow() {
	tt.Fail()
	tt.Stopped = true // TODO: stop execution for real (exit goroutine)
}

// Failed, as in testing.T, reports whether the function has failed.
func (tt *TestTester) Failed() bool {
	return tt.failed
}

// Fatal, as in testing.T, is equivalent to Log followed by FailNow.
func (tt *TestTester) Fatal(args ...interface{}) {
	tt.Log(args...)
	tt.FailNow()
}

// Fatalf, as in testing.T, is equivalent to Logf followed by FailNow.
func (tt *TestTester) Fatalf(format string, args ...interface{}) {
	tt.Logf(format, args...)
	tt.FailNow()
}

// Log, as in testing.T, records a log event a la Println.
func (tt *TestTester) Log(args ...interface{}) {
	// STUPID HACK WARNING: this may not work.
	if len(args) > 0 {
		f := make([]string, len(args))
		for i, _ := range args {
			f[i] = "%v"
		}
		format := strings.Join(f, " ")
		tt.Logf(format, args...)
	}
}

// Logf, as in testing.T, records a log event a la Printf.
func (tt *TestTester) Logf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	tt.Logs = append(tt.Logs, msg)
}

// Skip, as in testing.T, is equivalent to Log followed by SkipNow.
func (tt *TestTester) Skip(args ...interface{}) {
	tt.Log(args...)
	tt.SkipNow()
}

// SkipNow, as in testing.T, marks the test as having been skipped and stops
// its execution.
func (tt *TestTester) SkipNow() {
	tt.skipped = true
	tt.Stopped = true // TODO: stop execution for real (exit goroutine)

}

// Skipf, as in testing.T, is equivalent to Logf followed by SkipNow.
func (tt *TestTester) Skipf(format string, args ...interface{}) {
	tt.Logf(format, args...)
	tt.SkipNow()
}

// Skipped, as in testing.T, reports whether the test was skipped.
func (tt *TestTester) Skipped() bool {
	return tt.skipped
}

// AssertPanicsWith fails with msg unless the function f panics with string
// exp.
func AssertPanicsWith(t TT, f func(), exp, msg string) {

	panicked := false
	got := ""
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
				got = fmt.Sprintf("%s", r)
			}
		}()
		f()
	}()

	// NOTE: for testability without extra goroutines we make sure there is
	// no posibility of the test continuing after a Fail.
	// Also note: we lean on assert here because its failure messages are
	// so nice. :-)
	if !panicked {
		assert.Fail(t, "Function did not panic.", msg)
		t.FailNow()
	} else if got != exp {
		errMsg := fmt.Sprintf(
			"Panic not as expected:\n  expected: %s\n    actual: %s",
			exp, got)
		assert.Fail(t, errMsg, msg)
	}

	// (In go testing, success is silent.)

}
