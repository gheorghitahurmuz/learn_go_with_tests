package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Gheorghita")
	want := "Hello, Gheorghita!"

	if got != want {
		t.Errorf("got: %q,  want: %q", got, want)
	}
}
