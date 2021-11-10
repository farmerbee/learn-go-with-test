package main

import "testing"

func TestRepeat(t *testing.T) {
	want := "aaaaa"
	got := Repeat("a")

	if want != got {
		t.Errorf("want:%q, got:%q\n", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
