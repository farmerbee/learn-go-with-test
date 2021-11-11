package main

import (
	"bytes"
	"testing"
)

func TestDi(t *testing.T) {
	var buf bytes.Buffer
	Greet(&buf, "golang")
	want := "hello golang"

	if want != buf.String() {
		t.Errorf("want:%s, got:%s\n", want, buf.String())
	}
}
