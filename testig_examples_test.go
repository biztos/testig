// testig_examples_test.go

package testig_test

import (
	"testing"

	"github.com/biztos/testig"
)

func SomeTestFunction(t testing.TB) {
	t.Fatal("oh no")
}

func Example() {

	t := testig.NewTestTester()
	t.Log("here")

}
