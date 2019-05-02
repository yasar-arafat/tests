package main

import "testing"

func TestHello(t *testing.T) {

	got := Hello("Yasar")
	want := "Hello, Yasar"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
