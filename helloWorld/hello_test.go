package helloWorld

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("saying hello to people in english", func(t *testing.T) {
		got, _ := Hello("English", "Nabin")
		want := "Hello, Nabin"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got, _ := Hello("Spanish", "Nabin")
		want := "Hola, Nabin"

		assertHelloMsg(t, got, want)
	})

	t.Run("say 'Hello, World' when no name provided", func(t *testing.T) {
		got, _ := Hello("English")
		want := "Hello, World"

		assertHelloMsg(t, got, want)
	})

	t.Run("say 'Hello, World' when no argument provided", func(t *testing.T) {
		got, _ := Hello()
		want := "Hello, World"

		assertHelloMsg(t, got, want)
	})

	t.Run("passing more than two arguments", func(t *testing.T) {
		got, err := Hello("spanish", "May", "seen")
		want := ""

		assertHelloMsg(t, got, want)
		assertError(t, err, errArgumentNumber)
	})

}

func assertHelloMsg(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Wanted an error but didn't got one")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleHello() {
	sayingHello, _ := Hello("spanish", "Nabin")
	fmt.Println(sayingHello)

	// Output: Hola, Nabin
}

func BenchmarkHello(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Hello("Spanish", "Nabin")
	}
}
