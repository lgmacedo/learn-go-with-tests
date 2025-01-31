package arrays_and_slices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	assertMultipleSlicesItemsSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("array with size 5", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers) // poderia usar numbers[:] para mandar pra SumSlice. é basicamente uma conversão de array pra slice
		want := 15
		assertArrayOrSliceItemsSum(t, got, want, numbers[:])
	})

	t.Run("slice of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := SumSlice(numbers)
		want := 6
		assertArrayOrSliceItemsSum(t, got, want, numbers)
	})

	t.Run("sums multiple slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertMultipleSlicesItemsSum(t, got, want)
	})

	t.Run("sums multiple slices option 2", func(t *testing.T) {
		got := SumAllOption2([]int{6, 9}, []int{4, 2, 0})
		want := []int{15, 6}
		assertMultipleSlicesItemsSum(t, got, want)
	})

	t.Run("sums only tails of multiple slices", func(t *testing.T) {
		got := SumAllTails([]int{4, 7, 9}, []int{1, 2})
		want := []int{16, 2}
		assertMultipleSlicesItemsSum(t, got, want)
	})

	t.Run("sums only tails of multiple slices with one slice being empty", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2})
		want := []int{0, 2}
		assertMultipleSlicesItemsSum(t, got, want)
	})
}

func assertArrayOrSliceItemsSum(t *testing.T, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
