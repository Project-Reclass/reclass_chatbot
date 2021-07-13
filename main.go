package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

type Text struct {
	UserID   int    `json:"id"`
	Content  string `json:"content"`
	Username string `json:"username"`
	Date     string `json:"created"`
}

// Not sure if it's ok to not return an error?
func GetLatestId() (int, bool, error) {
	// Use the imported net/http to 'get' request and read from the API
	contributorsAPI, err := http.Get("http://192.168.49.2:30660/api/messages")
	if err != nil {
		return 0, false, err
	}
	defer contributorsAPI.Body.Close()
	body, err := ioutil.ReadAll(contributorsAPI.Body)
	if err != nil {
		return 0, false, err
	}

	// Unmarshal the body string into structs using Cont
	var tex []Text
	if marshalError := json.Unmarshal(body, &tex); marshalError != nil {
		return 0, false, marshalError
	}

	lastId := 0
	for i := 0; i < len(tex); i++ {
		thisId := tex[i].UserID
		if thisId > lastId {
			lastId = thisId
		}
	}
	return lastId, true, nil
}

func MainOutput(out io.Writer) {

	//Creating properly formatted time
	current := time.Now()
	currentFormat := current.Format(time.RubyDate)
	datept1 := regexp.MustCompile(`[a-zA-Z]+\s[a-zA-Z]+\s[0-9]*`)
	datept2 := regexp.MustCompile(`[0-9]*:[0-9]*:[0-9]*`)
	datept3 := regexp.MustCompile(`-[0-9]*`)
	finalDate := fmt.Sprintf(
		"%s %d %s GMT%s (Central Daylight Time)",
		datept1.FindStringSubmatch(currentFormat)[0],
		time.Now().Year(),
		datept2.FindStringSubmatch(currentFormat)[0],
		datept3.FindStringSubmatch(currentFormat)[0])

	// Putting together post via Text struct
	// Catch exception here! If valid is false but error is true
	prevId, valid, err := GetLatestId()
	if err != nil {
		fmt.Fprintln(out, err)
	} else if !valid {
		fmt.Fprintln(out, "Valid is false")
	}
	currentId := prevId + 1
	text := Text{currentId, finalDate, "Reclass Bot", finalDate}

	// Marshal the string to json
	jsonReq, err := json.Marshal(text)
	if err != nil {
		fmt.Fprintln(out, err)
	}

	// Post the json to the API
	resp, err := http.Post("http://192.168.49.2:30660/api/messages", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Fprintln(out, err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}

func main() {
	MainOutput(os.Stdout)
}
