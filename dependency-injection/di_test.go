package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gopher")

	got := buffer.String()
	want := "Hello, Gopher"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
