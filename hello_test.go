package main

import "testing"

func TestHello(t *testing.T) {

	/*testing.TB is an interface that *testing.T and *testing.B both satisfy,
	so you can call helper functions from a test, or a benchmark.*/
	assertMessage := func(t testing.TB, got, want string) {
		/*
			t.Helper() tells the test suite this variable function is a helper,
			so when it fails the reported line number will be
			in the function that called it, not inside the test helper
			Example:
			1. t.Helper() commented out => FAIL hello_test.go:20:
			2. t.Helper() not commented out => FAIL hello_test.go:35:
		*/
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	//using subtests here
	t.Run("customarily hello world", func(t *testing.T) {
		got := hello("welt")
		want := "Hello welt"
		assertMessage(t, got, want)
	})

	t.Run("say hello to people", func(t *testing.T) {

		got := hello("Ann")
		want := "Hello Ann"
		assertMessage(t, got, want)
	})

	t.Run("use default 'world' if no parameter is supplied", func(t *testing.T) {
		got := hello("")
		want := "Hello world"
		assertMessage(t, got, want)
	})
}
