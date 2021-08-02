package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"rsc.io/quote"
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
func CreatePost(prevId int, username string, message string) ([]byte, error) {
	//Creating properly formatted time
	current := time.Now().Format("Mon Jan 2 2006 15:04:05 GMT-0700 (MST)")

	text := Text{prevId + 1, message, username, current}

	// Marshal the string to json
	jsonReq, err := json.Marshal(text)
	if err != nil {
		return nil, err
	}
	return jsonReq, nil
}

func MainOutput(out io.Writer, username string, message string) {
	// Use the imported net/http to 'get' request and read from the API
	botAPI, err := http.Get("http://192.168.49.2:31592/api/messages")
	if err != nil {
		panic(err)
	}
	defer botAPI.Body.Close()
	body, err := ioutil.ReadAll(botAPI.Body)
	if err != nil {
		panic(err)
	}
	prevId, err := GetLatestId(body)
	if err != nil {
		panic(err)
	}
	// Create the post with the correct ID
	jsonpost, err := CreatePost(prevId, username, message)
	if err != nil {
		panic(err)
	}
	// POST the json to the API
	resp, err := http.Post(
		"http://192.168.49.2:31592/api/messages",
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

	fmt.Println("The program has started!")

	current := time.Now().Format("Mon Jan 2 2006 15:04:05 GMT-0700 (MST)")
	var (
		username = flag.String("username", "Reclass Robot", "a string")
		message  = flag.String("message", current, "a string")
		interval = flag.Int("interval", 3, "an int")
		random   = flag.Bool("random", false, "a boolean")
	)
	flag.Parse()

	switch *username {
	case "Tay":
		*message = quote.Hello()
		*random = false
		*interval = 3
	case "Kunal":
		*message = quote.Glass()
		*random = false
		*interval = 5
	case "Theo":
		*message = quote.Go()
		*random = false
		*interval = 7
	case "Scott":
		*message = quote.Opt()
		*random = false
		*interval = 9
	default:
		fmt.Println("Custom bot in use...")
	}

	// Repeatedly call the MainOutput() function
	var timer *time.Timer
	upperRange := *interval
	rand.Seed(time.Now().UnixNano())
	for {
		if *random {
			*interval = rand.Intn(upperRange)
			fmt.Println("delaying " + fmt.Sprint(*interval) + " seconds")
		}
		timer = time.NewTimer(time.Duration(*interval) * time.Second)
		<-timer.C
		MainOutput(os.Stdout, *username, *message)
	}
}
