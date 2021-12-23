package main

import (
	"log"
	"net/http"
)

// small webserver using go native packages
func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// this will start the server
	log.Fatal(http.ListenAndServe(":5000", nil))
}
