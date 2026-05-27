package main

import "net/http"

func handlerError(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, 400, msg)
	}
}
