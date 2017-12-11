// See Page 880.
//!+

//!+test
package charcount

import (
	"reflect"
	"testing"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input string
		want  map[rune]int
	}{
		{"mama",
			map[rune]int{
				'm': 2,
				'a': 2,
			},
		},
	}

	for _, test := range tests {
		runes := CharCount(test.input)
		if !reflect.DeepEqual(runes, test.want) {
			t.Errorf("%q runes, got %v, want %v", test.input, runes, test.want)
		}
	}
}

//!+bench
