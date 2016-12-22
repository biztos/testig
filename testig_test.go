// testig_test.go

package testig_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/biztos/testig"
)

func Test_NewTestTester(t *testing.T) {

	// Exercised elsewhere too of course, but let's be thorough.

	assert := assert.New(t)

	tt := testig.NewTestTester()
	assert.IsType(&testig.TestTester{}, tt, "has expected type")
	assert.Equal([]string{}, tt.Logs, "Logs empty at start")
	assert.False(tt.Stopped, "Stopped is false at start")
	assert.False(tt.Skipped(), "Skipped returns false")
	assert.False(tt.Failed(), "Failed returns false")

}

func Test_TestTester_Error(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.Error("foo", []int{1, 2, 3}, true)
	exp := []string{"foo [1 2 3] true"}
	assert.Equal(exp, tt.Logs, "logged as expected")
	assert.True(tt.Failed(), "Failed returns true")
	assert.False(tt.Skipped(), "Skipped returns false")
	assert.False(tt.Stopped, "TestTester not Stopped")

}

func Test_TestTester_Errorf(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.Errorf("foo %s %0.2f %t", "bar", 0.12345, true)
	exp := []string{"foo bar 0.12 true"}
	assert.Equal(exp, tt.Logs, "logged as expected")
	assert.True(tt.Failed(), "Failed returns true")
	assert.False(tt.Skipped(), "Skipped returns false")
	assert.False(tt.Stopped, "TestTester not Stopped")

}

func Test_TestTester_Fail(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.Fail()
	exp := []string{}
	assert.Equal(exp, tt.Logs, "(nothing) logged as expected")
	assert.True(tt.Failed(), "Failed returns true")
	assert.False(tt.Skipped(), "Skipped returns false")
	assert.False(tt.Stopped, "TestTester not Stopped")

}

func Test_TestTester_FailNow(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.FailNow()
	exp := []string{}
	assert.Equal(exp, tt.Logs, "(nothing) logged as expected")
	assert.True(tt.Failed(), "Failed returns true")
	assert.False(tt.Skipped(), "Skipped returns false")
	assert.True(tt.Stopped, "TestTester Stopped")

}

func Test_TestTester_Fatal(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.Fatal("foo", []int{1, 2, 3}, true)
	exp := []string{"foo [1 2 3] true"}
	assert.Equal(exp, tt.Logs, "logged as expected")
	assert.True(tt.Failed(), "Failed returns true")
	assert.False(tt.Skipped(), "Skipped returns false")
	assert.True(tt.Stopped, "TestTester Stopped")

}

func Test_TestTester_Fatalf(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.Fatalf("foo %s %0.2f %t", "bar", 0.12345, true)
	exp := []string{"foo bar 0.12 true"}
	assert.Equal(exp, tt.Logs, "logged as expected")
	assert.True(tt.Failed(), "Failed returns true")
	assert.False(tt.Skipped(), "Skipped returns false")
	assert.True(tt.Stopped, "TestTester Stopped")

}

func Test_TestTester_Log(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.Log("foo", []int{1, 2, 3}, true)
	exp := []string{"foo [1 2 3] true"}
	assert.Equal(exp, tt.Logs, "logged as expected")
	assert.False(tt.Failed(), "Failed returns false")
	assert.False(tt.Skipped(), "Skipped returns false")
	assert.False(tt.Stopped, "TestTester not Stopped")

}

func Test_TestTester_Log_Empty(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.Log()
	exp := []string{""}
	assert.Equal(exp, tt.Logs, "logged one empty string as expected")
	assert.False(tt.Failed(), "Failed returns false")
	assert.False(tt.Skipped(), "Skipped returns false")
	assert.False(tt.Stopped, "TestTester not Stopped")

}

func Test_TestTester_Logf(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.Logf("foo %s %0.2f %t", "bar", 0.12345, true)
	exp := []string{"foo bar 0.12 true"}
	assert.Equal(exp, tt.Logs, "logged as expected")
	assert.False(tt.Failed(), "Failed returns false")
	assert.False(tt.Skipped(), "Skipped returns false")
	assert.False(tt.Stopped, "TestTester not Stopped")

}

func Test_TestTester_Skip(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.Skip("foo", []int{1, 2, 3}, true)
	exp := []string{"foo [1 2 3] true"}
	assert.Equal(exp, tt.Logs, "logged as expected")
	assert.False(tt.Failed(), "Failed returns false")
	assert.True(tt.Skipped(), "Skipped returns true")
	assert.True(tt.Stopped, "TestTester Stopped")

}

func Test_TestTester_SkipNow(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.SkipNow()
	exp := []string{}
	assert.Equal(exp, tt.Logs, "(nothing) logged as expected")
	assert.False(tt.Failed(), "Failed returns false")
	assert.True(tt.Skipped(), "Skipped returns true")
	assert.True(tt.Stopped, "TestTester Stopped")

}

func Test_TestTester_Skipf(t *testing.T) {

	assert := assert.New(t)

	tt := testig.NewTestTester()
	tt.Skipf("foo %s %0.2f %t", "bar", 0.12345, true)
	exp := []string{"foo bar 0.12 true"}
	assert.Equal(exp, tt.Logs, "logged as expected")
	assert.False(tt.Failed(), "Failed returns false")
	assert.True(tt.Skipped(), "Skipped returns true")
	assert.True(tt.Stopped, "TestTester Stopped")

}
