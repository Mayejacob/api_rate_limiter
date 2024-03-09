package main

import (
	"encoding/json"
	"log"
	"net/http"

	tollbooth "github.com/didip/tollbooth/v7"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	message := Message{
		Status: "successful",
		Body:   "Hi! you've reached the Api, how may i help you?",
	}

	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

func main() {
	message := Message{
		Status: "failed",
		Body:   "The Api is at capacity, try again later.",
	}
	jsonMessage, _ := json.Marshal(message)
	tollboothLimiter := tollbooth.NewLimiter(1, nil)
	tollboothLimiter.SetMessageContentType("application/json")
	tollboothLimiter.SetMessage(string(jsonMessage))
	http.Handle("/ping", tollbooth.LimitFuncHandler(tollboothLimiter, endpointHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There was an error listening on port :8080", err)
	}
}
