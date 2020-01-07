package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Robert")
	want := "Hello, Robert"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
