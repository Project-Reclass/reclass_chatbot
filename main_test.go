package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestOutput(t *testing.T) {
	buf := &bytes.Buffer{}

	MainOutput(buf)

	got := buf.String()
	want := fmt.Sprintf("{%q:null,%q:true}\n", "error", "status")
	if got != want {
		t.Fatalf("Got: %s - Want: %s\n", got, want)
	}

	prevId, err := GetLatestId()
	if err != nil {
		t.Fatalf(err.Error())
	} else if prevId < 0 {
		t.Fatalf("Invalid: GetLatestId() returns a negative ID")
	}

}
