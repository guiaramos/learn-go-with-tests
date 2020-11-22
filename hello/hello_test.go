package main

import "testing"

func TestHello(t *testing.T) {

	assetCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Gui", "")
		want := "Hello, Gui"

		assetCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assetCorrectMessage(t, got, want)
	})

	t.Run("in portuguese", func(t *testing.T) {
		got := Hello("Gui", "Portuguese")
		want := "Ola, Gui"

		assetCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Gui", "French")
		want := "Bonjour, Gui"

		assetCorrectMessage(t, got, want)
	})
}


