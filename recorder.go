// recorder.go -- testing output recorders &c.

package testig

import (
	"bufio"
	"bytes"
	"fmt"
	"time"
)

// OutputRecorder allows recording and inspection of output to the normal
// channels Stdout and Stderr, as well as capture of the exit code.
type OutputRecorder struct {
	Stdout   *bufio.Writer
	Stderr   *bufio.Writer
	Exited   bool
	ExitCode int
	ExitTime time.Time

	outBuf *bytes.Buffer
	errBuf *bytes.Buffer
}

// NewOutputRecorder returns an initialied OutputRecorder ready for use.
func NewOutputRecorder() *OutputRecorder {

	var outBytes bytes.Buffer
	var errBytes bytes.Buffer

	return &OutputRecorder{
		Stdout:   bufio.NewWriter(&outBytes),
		Stderr:   bufio.NewWriter(&errBytes),
		outBuf:   &outBytes,
		errBuf:   &errBytes,
		ExitCode: -1,
	}
}

// Exit is a function suitable for overriding os.Exit.  If Exit is called more
// than once it panics: in order for exiting functions to be reasonably
// testable, they must not assume their exit calls actually terminate the
// program.
func (r *OutputRecorder) Exit(code int) {
	if r.Exited {
		panic("Exit called more than once; last was: " + r.ExitString())
	}
	r.Exited = true
	r.ExitCode = code
	r.ExitTime = time.Now()
}

// ExitString stringifies the exit status.
func (r *OutputRecorder) ExitString() string {

	if !r.Exited {
		return "did not exit"
	}
	return fmt.Sprintf("exit code %d at %v", r.ExitCode, r.ExitTime)
}

// StdoutString returns a string of all written to standard output so far.
func (r *OutputRecorder) StdoutString() string {
	r.Stdout.Flush()
	return r.outBuf.String()
}

// StderrString returns a string of all written to standard error so far.
func (r *OutputRecorder) StderrString() string {
	r.Stderr.Flush()
	return r.errBuf.String()
}
