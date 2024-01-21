package main

import "testing"

func TestHelloWorld(t *testing.T) {
	want := Hello("Mitch")
	got := "Hello, Mitch!"

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}

}
