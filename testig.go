// testig.go - the basics, rigging, self-testability, etc.

// Package testig provides helpers for testing Go programs, including itself.
// The spelling of the package name is intentional.
//
// LIMITATIONS
//
// There is currently no way for a TestTester to terminate a test function
// under test.  At the moment it doesn't seem worth the extra complexity of
// running tests in a separate goroutine, as the workaround also happens to
// be a useful practice: helper functions must not have any ability to
// continue after a Fail or Skip.
package testig

import (
	"fmt"
	"strings"
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

// Error mirrors the same-named function in testing.T: it is equivalent to Log
// followed by Fail.
func (tt *TestTester) Error(args ...interface{}) {
	tt.Log(args...)
	tt.Fail()

}

// Errorf mirrors the same-named function in testing.T: it is equivalent to
// Logf followed by Fail.
func (tt *TestTester) Errorf(format string, args ...interface{}) {
	tt.Logf(format, args...)
	tt.Fail()
}

// Fail mirrors the same-named function in testing.T: it marks the function as
// having failed but continues execution.
func (tt *TestTester) Fail() {
	tt.failed = true
}

// FailNow mirrors the same-named function in testing.T: it marks the
// function as having failed and sets the TestTester's Stopped property to
// true.
func (tt *TestTester) FailNow() {
	tt.Fail()
	tt.Stopped = true // TODO: stop execution for real (exit goroutine)
}

// Failed mirrors the same-named function in testing.T: it reports whether the
// function has failed.
func (tt *TestTester) Failed() bool {
	return tt.failed
}

// Fatal mirrors the same-named function in testing.T: it is equivalent to Log
// followed by FailNow.
func (tt *TestTester) Fatal(args ...interface{}) {
	tt.Log(args...)
	tt.FailNow()
}

// Fatalf mirrors the same-named function in testing.T: it is equivalent to
// Logf followed by FailNow.
func (tt *TestTester) Fatalf(format string, args ...interface{}) {
	tt.Logf(format, args...)
	tt.FailNow()
}

// Log mirrors the same-named function in testing.T: it records a log event a
// la Println.
func (tt *TestTester) Log(args ...interface{}) {
	// STUPID HACK WARNING: this may not work.
	if len(args) > 0 {
		f := make([]string, len(args))
		for i := range args {
			f[i] = "%v"
		}
		format := strings.Join(f, " ")
		tt.Logf(format, args...)
	} else {
		tt.Logf("")
	}
}

// Logf mirrors the same-named function in testing.T: it records a log event
// a la Printf.
func (tt *TestTester) Logf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	tt.Logs = append(tt.Logs, msg)
}

// Skip mirrors the same-named function in testing.T: it is equivalent to Log
// followed by SkipNow.
func (tt *TestTester) Skip(args ...interface{}) {
	tt.Log(args...)
	tt.SkipNow()
}

// SkipNow mirrors the same-named function in testing.T: it marks the test as
// having been skipped and sets the TestTester's Stopped property to true.
func (tt *TestTester) SkipNow() {
	tt.skipped = true
	tt.Stopped = true // TODO: stop execution for real (exit goroutine)

}

// Skipf mirrors the same-named function in testing.T: it is equivalent to
// Logf followed by SkipNow.
func (tt *TestTester) Skipf(format string, args ...interface{}) {
	tt.Logf(format, args...)
	tt.SkipNow()
}

// Skipped mirrors the same-named function in testing.T: it reports whether
// the test was skipped.
func (tt *TestTester) Skipped() bool {
	return tt.skipped
}
