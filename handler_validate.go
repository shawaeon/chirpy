package main

import (
	"encoding/json"
	"net/http"
)

// Chirp length validation
func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	const maxChirpLength = 140
	type parameters struct {
		Body string `json:"body"`
	}
	type validResponse struct {
		Valid bool `json:"valid"`
	}

	params := parameters{}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding JSON", err)
		return
	}

	if len(params.Body) > maxChirpLength {		
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)		
		return
	}	
	
	respondWithJSON(w, http.StatusOK, validResponse{
		Valid: true,
	})
}