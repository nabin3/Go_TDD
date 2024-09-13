package array_slice

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	expectedSummation := Sum(numbers)
	resultedSummation := 6

	if expectedSummation != resultedSummation {
		t.Errorf("expected %d got %d, given %v", expectedSummation, resultedSummation, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3, 5}, []int{2, 8})
	want := []int{9, 10}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sum of some slice's tails", func(t *testing.T) {
		got := SumAllTails([]int{2, 5, 8}, []int{3, 4, 5})
		want := []int{13, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
