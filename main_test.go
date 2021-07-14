package main

import (
	"encoding/json"
	"testing"
)

func TestOutput(t *testing.T) {

	// Check if CreatePost creates post with correct ID
	// CreatePost(prevId int) ([]byte, error)
	post, err := CreatePost(10, "Test Robot", "Hello World")
	if err != nil {
		t.Fatal(err)
	}
	var tex Text
	if marshalError := json.Unmarshal(post, &tex); marshalError != nil {
		t.Fatal(marshalError)
	} else if tex.UserID != 11 {
		t.Fatalf("Post ID should be 11, param + 1, instead %d,", tex.UserID)
	}
}
