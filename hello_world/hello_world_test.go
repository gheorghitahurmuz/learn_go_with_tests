package main

import "testing"

func TestHello(t *testing.T) {
	// these are subtests
	// they are useful to group tests aroung a "thing" and describing different scenarios
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gheorghita")
		want := "Hello, Gheorghita!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hell, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// this tell to test suite that this method is a helper
	// it help to track down problems easier (show the correct line from test)
	t.Helper()
	if got != want {
		t.Errorf("got: %q,  want: %q", got, want)
	}
}
