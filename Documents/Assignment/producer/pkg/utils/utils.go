package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseBody(request *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(request.Body)

	if err == nil {
		err = json.Unmarshal(body, &x)

		if err != nil {
			log.Print("Error while parsing body:", err)
			return
		}

	} else {
		return
	}
	log.Println("Successfully parsed body:", body)
}
