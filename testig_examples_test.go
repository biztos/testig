// testig_examples_test.go

package testig_test

import (
	"fmt"

	"github.com/biztos/testig"
)

var DETERMINISTIC_FAILURE = true

func Example() {

	// This will of course also accept a *testing.T as its argument.
	AwesomeTestingHelper := func(t testig.TT) {
		if DETERMINISTIC_FAILURE {
			t.Fatal("uh-oh spaghettio!")
		} else {
			t.Log("Things are looking up!")
		}
	}

	tt := testig.NewTestTester()
	AwesomeTestingHelper(tt)

	if tt.Stopped {
		fmt.Println("Failed:", tt.Failed())
		fmt.Println("Skipped:", tt.Skipped())
		fmt.Println(tt.Logs)
	}
	// Output:
	// Failed: true
	// Skipped: false
	// [uh-oh spaghettio!]

}

func ExampleAssertPanicsWith() {

	dontPanic := func() { return }
	doPanic := func() { panic("oh no") }

	tt := testig.NewTestTester()

	testig.AssertPanicsWith(tt, doPanic, "oh no", "got a scary panic")
	fmt.Println("Failed:", tt.Failed())

	testig.AssertPanicsWith(tt, dontPanic, "any panic", "got another")
	fmt.Println("Failed:", tt.Failed())

	// Output:
	// Failed: false
	// Failed: true
}
