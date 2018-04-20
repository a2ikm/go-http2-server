package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		for {
			fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
			flusher.Flush()
			time.Sleep(2 * time.Second)
		}
	})

	log.Print("Starting at :3000")
	err := http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
