package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetResponse(data interface{}, w http.ResponseWriter, status int) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var response = Response{
		Data: data,
	}

	message, _ := json.Marshal(response)
	w.WriteHeader(status)
	w.Write(message)
}

func GetError(err error, w http.ResponseWriter, status int) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	log.Println(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   status,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
