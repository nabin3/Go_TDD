package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 4)
	expected := 6

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func Example() {
	sum := Add(5, 7)
	fmt.Println(sum)

	// Output: 12
}
