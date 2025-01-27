package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("2 + 2 equals 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(2, 7)
	fmt.Println(sum)
	// Output: 9
}

func assertCorrectMessage(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", sum, expected)
	}
}
