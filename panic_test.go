// panic_test.go

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
		assert.Regexp("Function did not panic", tt.Logs[0],
			"expected stuff in logs")

		assert.Regexp("panicky func panicked", tt.Logs[0],
			"...including our test name")
	}

}

func Test_AssertPanicsWith_OtherPanic(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()

	panicky := func() {
		panic("not so fast")
	}

	testig.AssertPanicsWith(tt, panicky, "uh-oh", "panicky func panicked")

	assert.True(tt.Failed(), "test failed")
	if assert.Equal(1, len(tt.Logs), "one thing logged") {
		assert.Regexp("Panic not as expected", tt.Logs[0],
			"expected stuff in logs")

		assert.Regexp("expected: uh-oh", tt.Logs[0],
			"...including the expected panic text")

		assert.Regexp("actual: not so fast", tt.Logs[0],
			"...including the actual panic text")

		assert.Regexp("panicky func panicked",
			tt.Logs[0], "...including our test name")
	}

}

func Test_AssertPanicsWith_OtherPanic_NoMessage(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()

	panicky := func() {
		panic("not so fast")
	}

	testig.AssertPanicsWith(tt, panicky, "uh-oh")

	assert.True(tt.Failed(), "test failed")
	if assert.Equal(1, len(tt.Logs), "one thing logged") {
		assert.Regexp("Panic not as expected", tt.Logs[0],
			"expected stuff in logs")

		assert.Regexp("expected: uh-oh", tt.Logs[0],
			"...including the expected panic text")

		assert.Regexp("actual: not so fast", tt.Logs[0],
			"...including the actual panic text")

	}

}

func Test_AssertPanicsWith_OtherPanic_FormattedMessage(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()

	panicky := func() {
		panic("not so fast")
	}

	testig.AssertPanicsWith(tt, panicky, "uh-oh", "panicked: %d", 999)

	assert.True(tt.Failed(), "test failed")
	if assert.Equal(1, len(tt.Logs), "one thing logged") {
		assert.Regexp("Panic not as expected", tt.Logs[0],
			"expected stuff in logs")

		assert.Regexp("expected: uh-oh", tt.Logs[0],
			"...including the expected panic text")

		assert.Regexp("actual: not so fast", tt.Logs[0],
			"...including the actual panic text")

		assert.Regexp("panicked: 999",
			tt.Logs[0], "...including our test name")

	}

}

func Test_AssertPanicsRegexp_Success(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()

	panicky := func() { panic("uh-oh") }

	testig.AssertPanicsRegexp(tt, panicky, "^uh-..", "panicky func panicked")

	assert.False(tt.Failed(), "test did not fail")
	assert.Equal([]string{}, tt.Logs, "nothing logged")
}

func Test_AssertPanicsRegexp_NoPanic(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()

	panicky := func() {
		// don't panic!
	}

	testig.AssertPanicsRegexp(tt, panicky, ".*", "panicky func panicked")

	assert.True(tt.Failed(), "test failed")
	if assert.Equal(1, len(tt.Logs), "one thing logged") {
		assert.Regexp("Function did not panic", tt.Logs[0],
			"expected stuff in logs")

		assert.Regexp("panicky func panicked", tt.Logs[0],
			"...including our test name")
	}

}

func Test_AssertPanicsRegexp_OtherPanic(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()

	panicky := func() {
		panic("not so fast")
	}

	testig.AssertPanicsRegexp(tt, panicky, "^uh-..", "panicky func panicked")

	assert.True(tt.Failed(), "test failed")
	if assert.Equal(1, len(tt.Logs), "one thing logged") {
		assert.Regexp("Panic not as expected", tt.Logs[0],
			"expected stuff in logs")

		assert.Regexp("expected: Regexp /\\^uh-../", tt.Logs[0],
			"...including the expected panic text")

		assert.Regexp("actual: not so fast", tt.Logs[0],
			"...including the actual panic text")

		assert.Regexp("panicky func panicked",
			tt.Logs[0], "...including our test name")
	}

}

func Test_AssertPanicsRegexp_OtherPanic_NoMessage(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()

	panicky := func() {
		panic("not so fast")
	}

	testig.AssertPanicsRegexp(tt, panicky, "matchme")

	assert.True(tt.Failed(), "test failed")
	if assert.Equal(1, len(tt.Logs), "one thing logged") {
		assert.Regexp("Panic not as expected", tt.Logs[0],
			"expected stuff in logs")

		assert.Regexp("expected: Regexp /matchme/", tt.Logs[0],
			"...including the expected panic text")

		assert.Regexp("actual: not so fast", tt.Logs[0],
			"...including the actual panic text")

	}

}

func Test_AssertPanicsRegexp_OtherPanic_FormattedMessage(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()

	panicky := func() {
		panic("not so fast")
	}

	testig.AssertPanicsRegexp(tt, panicky, "^uh-..", "panicked: %d", 999)

	assert.True(tt.Failed(), "test failed")
	if assert.Equal(1, len(tt.Logs), "one thing logged") {
		assert.Regexp("Panic not as expected", tt.Logs[0],
			"expected stuff in logs")

		assert.Regexp("expected: Regexp /\\^uh-../", tt.Logs[0],
			"...including the expected panic text")

		assert.Regexp("actual: not so fast", tt.Logs[0],
			"...including the actual panic text")

		assert.Regexp("panicked: 999",
			tt.Logs[0], "...including our test name")

	}

}
