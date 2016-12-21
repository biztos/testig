// testig_test.go

package testig_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/biztos/testig"
)

func Test_AssertPanicsWith_Success(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()

	panicky := func() { panic("uh-oh") }

	testig.AssertPanicsWith(tt, panicky, "uh-oh", "panicky func panicked")

	assert.False(tt.Failed(), "test did not fail")
	assert.Equal([]string{}, tt.Logs, "nothing logged")
}

func Test_AssertPanicsWith_NoPanic(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()

	panicky := func() {
		// don't panic!
	}

	testig.AssertPanicsWith(tt, panicky, "uh-oh", "panicky func panicked")

	assert.True(tt.Failed(), "test failed")
	if assert.Equal(1, len(tt.Logs), "one thing logged") {
		assert.Regexp("Function did not panic",
			tt.Logs[0], "expected stuff in logs")
		assert.Regexp("panicky func panicked",
			tt.Logs[0], "...including our test name")
	}

}
