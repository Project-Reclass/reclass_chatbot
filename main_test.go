package main

import (
	"encoding/json"
	"testing"
)

func TestOutput(t *testing.T) {

	// Check if CreatePost creates post with correct ID
	// CreatePost(prevId int) ([]byte, error)
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

	// Having trouble with this one. I wanted to pass in the post I just created since it's []byte
	// but it results in a json unmarshal error. I think it's because the function expects
	// something like [{...},{...}] when 'post' is just {...}
	// GetLatestId(botAPI []byte]) (int, error)
	/*
		lastId, err := GetLatestId(post)
		if err != nil {
			t.Fatal(err)
		} else if lastId != 11 {
			t.Fatalf("Last ID should be 11, is %d", lastId)
		}
	*/

}
