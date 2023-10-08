package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseBody(request *http.Request, x interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&x)

	if err != nil {
		log.Print("Error while parsing body:", err)
		return
	}

	log.Println("Successfully parsed body")
}
