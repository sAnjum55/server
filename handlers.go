package main

import "net/http"

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

func handleError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "something went wrong")
}
