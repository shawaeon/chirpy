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
	mux.Handle("/app/", cfg.middlewareMetricsInc(
		http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("/healthz", handlerReadiness)
	mux.HandleFunc("/metrics", cfg.handlerMetrics)
	mux.HandleFunc("/reset", cfg.handlerResetMetrics)

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

// Middleware to keep count of requests
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})	
}

