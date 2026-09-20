package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to a person", func(t *testing.T) {
		got := Hello("Fulan", "")
		want := "Hello, Fulan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when `who` is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Budi in Indonesia", func(t *testing.T) {
		got := Hello("Budi", "id")
		want := "Halo, Budi"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Budi in Arabic", func(t *testing.T) {
		got := Hello("Budi", "ar")
		want := "Ahlan, Budi"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
