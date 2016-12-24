// recorder_test.go

package testig_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/biztos/testig"
)

func Test_NewOutputRecorder(t *testing.T) {

	assert := assert.New(t)

	r := testig.NewOutputRecorder()
	assert.NotNil(r.Stdout, "Stdout initialized")
	assert.NotNil(r.Stderr, "Stdout initialized")
	assert.Equal(-1, r.ExitCode, "ExitCode set to -1")

}

func Test_OutputRecorder_ExitString(t *testing.T) {

	assert := assert.New(t)

	r := testig.NewOutputRecorder()

	assert.Equal("did not exit", r.ExitString(),
		"pre-exit state stringified as expected")

	r.Exit(123)

	// NOTE: not checking the timestamp per se, as that could differ based
	// on the tester's OS and settings.  TODO: make that work!
	assert.Regexp("^exit code 123 at \\S+", r.ExitString(),
		"post-exit state stringified")
}

func Test_OutputRecorder_Exit_PanicsOnSecond(t *testing.T) {

	r := testig.NewOutputRecorder()
	r.Exit(123)

	// A fine use for our own panic checkers!
	testig.AssertPanicsRegexp(t, func() { r.Exit(321) },
		"^Exit called more than once; last was: exit code 123 at \\S+")
}

func Test_OutputRecorder_StdoutString(t *testing.T) {

	assert := assert.New(t)

	r := testig.NewOutputRecorder()

	fmt.Fprintln(r.Stdout, "first line")
	fmt.Fprint(r.Stdout, "second", " ", "line")
	fmt.Fprintln(r.Stdout)

	assert.Equal("first line\nsecond line\n", r.StdoutString(),
		"StdoutString returns buffer string")
}

func Test_OutputRecorder_StderrString(t *testing.T) {

	assert := assert.New(t)

	r := testig.NewOutputRecorder()

	fmt.Fprintln(r.Stderr, "first line")
	fmt.Fprint(r.Stderr, "second", " ", "line")
	fmt.Fprintln(r.Stderr)

	assert.Equal("first line\nsecond line\n", r.StderrString(),
		"StderrString returns buffer string")
}
