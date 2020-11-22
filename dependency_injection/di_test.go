package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gui")

	got := buffer.String()
	want := "Hello, Gui"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
