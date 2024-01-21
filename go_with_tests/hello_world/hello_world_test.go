package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		want := Hello("Mitch")
		got := "Hello, Mitch!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Default to 'Hello, World!' when no name is supplied", func(t *testing.T) {

		want := Hello("")
		got := "Hello, World!"

		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
