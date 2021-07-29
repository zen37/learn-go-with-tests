package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("welt")
	want := "hello, welt"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello2(t *testing.T) {
	got := hello("Chris")
	want := "hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
