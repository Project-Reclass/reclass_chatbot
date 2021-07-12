package main

import (
	"bytes"
	"testing"
)

func TestOutput(t *testing.T) {
	buf := &bytes.Buffer{}

	MainOutput(buf)

	got := buf.String()
	want := "Hello World!\n"
	if got != want {
		t.Fatalf("Got: %s - Want: %s\n", got, want)
	}
}
