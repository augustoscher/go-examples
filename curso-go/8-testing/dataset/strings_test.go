package dataset

import (
	"strings"
	"testing"
)

const msg = "%s (part: %s) - Index: expected (%d) <> was (%d)."

func TestIndex(t *testing.T) {
	//array of structs containing test cases
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		//Testing word is in 0 indes of "Testing is great" string
		{"Testing is greate", "Testing", 0},
		//Empty string is in 0 index to
		{"", "", 0},
		//hey is not present in first string. -1 expected
		{"Hey", "hey", -1},
		//how word is in 4 index of given text
		{"Hi, how are you?", "how", 4},
		{"augusto", "g", 2},
	}

	for _, test := range tests {
		t.Logf("Test: %v", test)
		actual := strings.Index(test.text, test.part)

		if actual != test.expected {
			t.Errorf(msg, test.text, test.part, test.expected, actual)
		}
	}
}
