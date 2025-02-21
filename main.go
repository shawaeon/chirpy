package main

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	cfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}
	
	mux := http.NewServeMux()
	
	fsHandler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))
	mux.Handle("/app/", cfg.middlewareMetricsInc(fsHandler))

	mux.HandleFunc("GET /healthz", handlerReadiness)
	mux.HandleFunc("GET /metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /reset", cfg.handlerResetMetrics)

	srv := &http.Server {
		Addr: ":" + port,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())	
}

type apiConfig struct {
	fileserverHits atomic.Int32
}
