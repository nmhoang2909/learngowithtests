package arraysandslices_test

import (
	"testing"

	. "example.com/go-test/arrays_and_slices"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d but want %d, given %v", got, want, numbers)
		}
	})
}
