package main

import (
	"encoding/json"
	"testing"
)

func TestOutput(t *testing.T) {
	/*
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
	*/

	// test
	// GetLatestId(botAPI []byte]) (int, error)
	// CreatePost(prevId int) ([]byte, error)

	// Check if CreatePost creates post with correct ID
	post, err := CreatePost(10)
	if err != nil {
		t.Fatal(err)
	}
	var tex Text
	if marshalError := json.Unmarshal(post, &tex); marshalError != nil {
		t.Fatal(marshalError)
	} else if tex.UserID != 11 {
		t.Fatalf("Post ID should be 11, param + 1, instead %d,", tex.UserID)
	}

	lastId, err := GetLatestId(post)
	if err != nil {
		t.Fatal(err)
	} else if lastId != 11 {
		t.Fatalf("Last ID should be 11, is %d", lastId)
	}

}
