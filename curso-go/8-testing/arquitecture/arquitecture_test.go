package arquitecture

import (
	"testing"
	"runtime"
)

const defaultError = "Expected: %v - Current was %v."

//Test method has the same signature: Test + NameOfFunction
func TestDependent(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Not working in amd64 arquitecture")
	}

	t.Fail()
}
