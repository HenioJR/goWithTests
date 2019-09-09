package main

import "testing"

func TestHello(t *testing.T) {

	//In Go you can declare functions inside other functions and assign them to variables. You can then call them, just like normal functions
	assertCorrectMesssage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bob", "")
		want := "Hello, Bob"

		assertCorrectMesssage(t, got, want)
	})

	t.Run("saying hello world when an empty string is suplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMesssage(t, got, want)
	})

	t.Run("saying hello world in Spanish", func(t *testing.T) {
		got := Hello("Paredes", "Spanish")
		want := "Hola, Paredes"

		assertCorrectMesssage(t, got, want)
	})

	t.Run("saying hello world in French", func(t *testing.T) {
		got := Hello("Benzema", "French")
		want := "Bonjour, Paredes"

		assertCorrectMesssage(t, got, want)
	})

}
