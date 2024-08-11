package helloWorld

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	checker := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("English", "Nabin")
		want := "Hello, Nabin"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string provided", func(t *testing.T) {
		got := Hello("English", "")
		want := "Hello, World"

		checker(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Spanish", "Nabin")
		want := "Hola, Nabin"

		checker(t, got, want)
	})
}

func ExampleHello() {
	fmt.Println(Hello("Spanish", "Nabin"))

	// Output: Hola, Nabin
}

func BenchmarkHello(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Hello("Spanish", "Nabin")
	}
}
