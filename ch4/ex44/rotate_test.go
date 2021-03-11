package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		input []int
		times int
		want  []int
	}{
		// rotate 2
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		// rotate 4
		{[]int{1, 2, 3, 4, 5}, 4, []int{5, 1, 2, 3, 4}},
		// rotate 5
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		// rotate 6
		{[]int{1, 2, 3, 4, 5}, 6, []int{2, 3, 4, 5, 1}},
		// rotate 7 same as rotate 2
		{[]int{1, 2, 3, 4, 5}, 7, []int{3, 4, 5, 1, 2}},
	}

	for _, test := range tests {
		rotate(test.input, test.times)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("got %v but want %v\n", test.input, test.want)
		}
	}
}
func TestRotateV1(t *testing.T) {
	tests := []struct {
		input []int
		times int
		want  []int
	}{
		// rotate 2
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		// rotate 4
		{[]int{1, 2, 3, 4, 5}, 4, []int{5, 1, 2, 3, 4}},
		// rotate 5
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		// rotate 6
		{[]int{1, 2, 3, 4, 5}, 6, []int{2, 3, 4, 5, 1}},
		// rotate 7 same as rotate 2
		{[]int{1, 2, 3, 4, 5}, 7, []int{3, 4, 5, 1, 2}},
	}

	for _, test := range tests {
		got := rotateV1(test.input, test.times)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v but want %v\n", got, test.want)
		}
	}
}
