package array

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{1, 2, 3})
	want := []int{6, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v' ", got, want)
	}

}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v' ", got, want)
		}
	}

	t.Run("make sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5})
		want := []int{5, 9, 14}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3, 4, 5})
		want := []int{0, 14}
		checkSums(t, got, want)
	})

}
