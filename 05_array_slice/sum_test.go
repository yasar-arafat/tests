package array

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection any size ", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("want '%d' got '%d' given '%v'", want, got, numbers)
		}
	})

}
