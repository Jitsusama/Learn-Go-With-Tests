package main

import (
	"jitsusama/lgwt/di/greet"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		greet.Greet(w, "world")
	}
	err := http.ListenAndServe(":5000", http.HandlerFunc(handler))
	if err != nil {
		log.Fatal(err)
	}
}
