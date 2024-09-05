package array_slice

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		expectedSummation := Sum(numbers)
		resultedSummation := 6

		if expectedSummation != resultedSummation {
			t.Errorf("expected %d got %d, given %v", expectedSummation, resultedSummation, numbers)
		}
	})
}
