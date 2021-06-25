package main

import (
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	log.Println("Listening on :8083...")
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatal(err)
	}

}
