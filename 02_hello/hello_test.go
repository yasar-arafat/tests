package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// it fails the line number reported will be in our
		// function call rather than inside our test helper.
		//try failing it and check error in line number carfully
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Yasar", "")
		want := "Hello, Yasar"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say, 'Hellow, world' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Yasar", "Spanish")
		want := "Hola, Yasar"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Yasar", "French")
		want := "Bonjour, Yasar"
		assertCorrectMessage(t, got, want)
	})

}
