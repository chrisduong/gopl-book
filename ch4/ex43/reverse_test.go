package reverse

import (
	"reflect"
	"testing"
)

// Exercise 4.3: Re write reverse to use an array pointer instead of a slice - Page 112.

// NOTE: This is for illustrating the relationshop between the Underlying Array and a slice by using Pointer. This meant that user to pass an Array (not a slice).

func TestReverse(t *testing.T) {
	tests := []struct {
		input [5]int
		want  [5]int
	}{
		{[5]int{9, 3, 1, 2, 7}, [5]int{7, 2, 1, 3, 9}},
	}

	for _, test := range tests {
		reverse(&test.input)

		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("got %v but want %v\n", test.input, test.want)
		}
	}
}
