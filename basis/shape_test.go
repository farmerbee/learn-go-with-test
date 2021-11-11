package main

import "testing"

func Assert(t *testing.T, want, got float64) {
	t.Helper()
	if want != got {
		t.Errorf("want:%v, got:%v\n", want, got)
	}
}
func TestPerimeter(t *testing.T) {
	got := Perimeter(10, 10)
	want := 40.0

	Assert(t, want, got)
}

func TestArea(t *testing.T) {
	got := Area(10, 10)
	want := 100.0

	Assert(t, want, got)
}
