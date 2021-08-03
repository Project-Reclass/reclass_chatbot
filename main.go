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
	Content  string `json:"content"`
	Username string `json:"username"`
	Date     string `json:"created"`
}

// Given the greatest existing ID in the API, this creates json for a new post
func CreatePost(username string, message string) ([]byte, error) {
	//Creating properly formatted time
	current := time.Now().Format("Mon Jan 2 2006 15:04:05 GMT-0700 (MST)")

	text := Text{message, username, current}

	// Marshal the string to json
	jsonReq, err := json.Marshal(text)
	if err != nil {
		return nil, err
	}
	return jsonReq, nil
}

func MainOutput(out io.Writer, username string, message string) {
	//testing
	apiurl := os.Getenv("CHATBACK_URL")
	if apiurl == "" {
		panic("You must set the API environment variable using 'export CHATBACK_URL={ chatback url }' or adding -e ... in docker run command")
	}

	// Use the imported net/http to 'get' request and read from the API
	botAPI, err := http.Get(apiurl)
	if err != nil {
		panic(err)
	}
	defer botAPI.Body.Close()

	// Create the post with the correct ID
	jsonpost, err := CreatePost(username, message)
	if err != nil {
		panic(err)
	}
	// POST the json to the API
	resp, err := http.Post(
		apiurl,
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
		preset   = flag.String("preset", "", "a string")
	)
	flag.Parse()

	switch *preset {
	case "Tay":
		fmt.Println("Overriding username, message, random with preset")
		*username = "Tay"
		*message = quote.Hello()
		*random = false
		*interval = 3
	case "Kunal":
		fmt.Println("Overriding username, message, random with preset")
		*username = "Kunal"
		*message = quote.Glass()
		*random = false
		*interval = 5
	case "Theo":
		fmt.Println("Overriding username, message, random with preset")
		*username = "Theo"
		*message = quote.Go()
		*random = false
		*interval = 7
	case "Scott":
		fmt.Println("Overriding username, message, random with preset")
		*username = "Scott"
		*message = quote.Opt()
		*random = false
		*interval = 9
	case "":
		fmt.Println("Custom bot in use (no presets)")
	default:
		fmt.Println("Invalid preset selected. No preset applied.")
		fmt.Println("Presets include 'Tay','Kunal','Theo','Scott'.")
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
