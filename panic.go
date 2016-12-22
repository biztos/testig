// panic.go -- testing panicky things

package testig

import (
	"fmt"

	"github.com/stretchr/testify/assert"
)

// AssertPanicsWith fails with msgAndArgs and stops test execution unless the
// function f panics with string exp.  It is safe to omit msgAndArgs.
func AssertPanicsWith(t TT, f func(), exp string, msgAndArgs ...interface{}) {

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
		assert.FailNow(t, "Function did not panic.", msgAndArgs...)
	} else if got != exp {
		errMsg := fmt.Sprintf(
			"Panic not as expected:\n  expected: %s\n    actual: %s",
			exp, got)
		assert.FailNow(t, errMsg, msgAndArgs...)
	}

	// (In go testing, success is silent.)

}
