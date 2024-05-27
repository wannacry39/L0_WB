package myServer

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi from /hi")
	})

	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
