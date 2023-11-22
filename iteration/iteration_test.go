package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Positive Numbers", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectString(t, repeated, expected)
	})
	t.Run("Run Zero Times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assertCorrectString(t, repeated, expected)
	})
	t.Run("Run Negitive Times", func(t *testing.T) {
		repeated := Repeat("a", -5)
		expected := ""

		assertCorrectString(t, repeated, expected)
	})
}
func assertCorrectString(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	output := Repeat("b", 5)
	fmt.Println(output)
	//Output: bbbbb
}
