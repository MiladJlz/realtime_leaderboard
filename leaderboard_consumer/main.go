package main

import (
	"log"
	"net/http"
)

func main() {

	ds, err := NewDataSender()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ws", ds.handleWS)
	http.ListenAndServe(":40000", nil)
}
