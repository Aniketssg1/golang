package main

import (
	"testing"
)

// Group the tests according to overall objective and then write subtests
func TestHello(t *testing.T) {
	// subtests for different scenarios
	t.Run("saying hello to Aniket", func(t *testing.T) {
		got := Hello("Aniket", "Spanish")
		want := "Hola, Aniket!!!"
		assertCorrectGreetings(t, got, want)
	})

	t.Run("saying hello to  ", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!!!"
		assertCorrectGreetings(t, got, want)
	})
}

func assertCorrectGreetings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
