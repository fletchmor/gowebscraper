package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 400, struct {
		ErrorResponse string
	}{ErrorResponse: "Something went wrong"})
}
