package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(want, got string) {
		t.Helper()
		if want != got {
			t.Errorf("want:%q, got:%q\n", want, got)
		}
	}

	t.Run("none nil name", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "hello Chris"
		assertMessage(want, got)
	})

	t.Run("nil name", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour world"
		assertMessage(want, got)
	})
	t.Run("nil name", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola world"
		assertMessage(want, got)
	})
}
