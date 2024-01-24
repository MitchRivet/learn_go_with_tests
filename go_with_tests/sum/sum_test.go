package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("It should sum a list of 5 numbers", func(t *testing.T) {

		// slice is similar to array, just don't give size at declaration
		numbers := []int{1, 2, 3, 4, 5}

		expect := 15
		got := Sum(numbers)

		if expect != got {
			t.Errorf("expected %d, got %d", expect, got)

		}

	})
	t.Run("Sum a list of any size", func(t *testing.T) {

		// slice is similar to array, just don't give size at declaration
		numbers := []int{1, 2, 3}

		expect := 6
		got := Sum(numbers)

		if expect != got {
			t.Errorf("expected %d, got %d", expect, got)

		}

	})

}

func TestSumAll(t *testing.T) {

	t.Run("Sum an arbitrary number of slices together", func(t *testing.T) {

		// slice is similar to array, just don't give size at declaration

		expect := []int{10, 3}
		got := SumAll([]int{1, 9}, []int{0, 3})

		// can't directly compare value of slices, same as most languages
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %d, got %d", expect, got)

		}

	})

}
