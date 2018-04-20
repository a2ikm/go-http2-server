package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	})

	log.Print("Starting at :3000")
	err := http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
