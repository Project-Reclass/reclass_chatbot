package main

//Start here: Do unit testing by outsourcing the logic and testing
//those helper functions individually. E.g. do post and get, then
//pass the result to helper functions. Then test those helper functions
//with sample get and post results boi.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Text struct {
	UserID   int    `json:"id"`
	Content  string `json:"content"`
	Username string `json:"username"`
	Date     string `json:"created"`
}

// Return the greatest existing post ID in API, regardless of order
func GetLatestId(body []byte) (int, error) {
	// Unmarshal the body string into structs using Cont
	var tex []Text
	if marshalError := json.Unmarshal(body, &tex); marshalError != nil {
		return 0, marshalError
	}

	// Grab the greatest of all post IDs
	lastId := 0
	for i := 0; i < len(tex); i++ {
		thisId := tex[i].UserID
		if thisId > lastId {
			lastId = thisId
		}
	}
	return lastId, nil
}

// Given the greatest existing ID in the API, this creates json for a new post
func CreatePost(prevId int) ([]byte, error) {
	//Creating properly formatted time
	current := time.Now()
	finalDate := current.Format("Mon Jan 2 2006 15:04:05 GMT-0700 (MST)")

	text := Text{prevId + 1, finalDate, "Reclass Bot", finalDate}

	// Marshal the string to json
	jsonReq, err := json.Marshal(text)
	if err != nil {
		return nil, err
	}
	return jsonReq, nil
}

func MainOutput(out io.Writer) {
	// Use the imported net/http to 'get' request and read from the API
	botAPI, err := http.Get("http://192.168.49.2:30660/api/messages")
	if err != nil {
		panic(err)
	}
	defer botAPI.Body.Close()
	body, err := ioutil.ReadAll(botAPI.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(body))
	prevId, err := GetLatestId(body)
	if err != nil {
		panic(err)
	}
	// Create the post with the correct ID
	jsonpost, err := CreatePost(prevId)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(jsonpost))
	// POST the json to the API
	resp, err := http.Post(
		"http://192.168.49.2:30660/api/messages",
		"application/json; charset=utf-8",
		bytes.NewBuffer(jsonpost))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	// (question) How to make this only print if there's an error?
	bodyString := string(bodyBytes)
	fmt.Fprintln(out, bodyString)

}

func main() {
	MainOutput(os.Stdout)
}
