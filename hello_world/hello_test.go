package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to the world", func (t *testing.T){
		got := Hello("Shubham")
		want := "hello, Shubham"
		correctionMassage(t, got, want)
	})

	t.Run ("say 'hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "hello, world"
		correctionMassage(t, got, want)
	})
}

func correctionMassage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}