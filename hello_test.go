package main

import "testing"

func TestHello(t *testing.T) {

	//using subtests here
	t.Run("customarily hello world", func(t *testing.T) {
		got := hello("welt")
		want := "Hello welt"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello to people", func(t *testing.T) {

		got := hello("welt")
		want := "Hello welt"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
