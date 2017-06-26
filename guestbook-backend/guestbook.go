package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Entry struct {
	TimestampNano int64  `json:"time,string"`
	Author        string `json:"author"`
	Message       string `json:"message"`
}

var guestBook []Entry

func guestBookHandler(response http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	if author := request.FormValue("author"); author != "" {
		guestBook = append(guestBook, Entry{time.Now().UnixNano(), request.FormValue("author"), request.FormValue("message")})
	}

	if guestBook != nil {
		jsonresponse, _ := json.Marshal(guestBook)
		fmt.Fprintf(response, "%s", jsonresponse)
	}
}

func main() {
	http.HandleFunc("/", guestBookHandler)
	http.ListenAndServe(":8080", nil)
}