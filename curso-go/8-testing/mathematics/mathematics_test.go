package mathematics

import "testing"

const defaultError = "Expected: %v - Current was %v."

//Test method has the same signature: Test + NameOfFunction
func TestAverage(t *testing.T) {
	t.Parallel()
	expected := 7.28
	value := Average(7.2, 9.9, 6.1, 5.9)

	if value != expected {
		t.Errorf(defaultError, expected, value)
	}
}
