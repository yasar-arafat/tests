package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Yasar")

	got := buffer.String()
	want := "Hello, Yasar"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
