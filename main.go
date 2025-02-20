package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	const port = "8080"
	
	mux := http.NewServeMux()

	srv := &http.Server {
		Addr: ":" + port,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
