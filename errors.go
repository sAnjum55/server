package main

import (
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5xx error")
	}
	type errorResponse struct {
		ErrMess string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		ErrMess: message,
	})

}
