package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponseItem struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Title  string `json:"title"`
}

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Errors []ErrorResponseItem `json:"errors"`
}

func SendResponse(w http.ResponseWriter, payload interface{}, statusCode int) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(statusCode)
	w.Write(response)
}

func SendErrorResponse(w http.ResponseWriter, errors ErrorResponseItem, statusCode int) {
	payload := ErrorResponse{
		Errors: []ErrorResponseItem{
			errors,
		},
	}

	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(statusCode)
	w.Write(response)
}
