package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Hits: %v\n", cfg.fileserverHits.Load())
}

func (cfg *apiConfig) handlerResetMetrics(w http.ResponseWriter, req *http.Request) {
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Hits reset to %v\n", cfg.fileserverHits.Load())
}