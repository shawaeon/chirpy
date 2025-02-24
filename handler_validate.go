package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

// Chirp length validation
func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	const maxChirpLength = 140
	type parameters struct {
		Body string `json:"body"`
	}
	type validResponse struct {
		CleanedBody string `json:"cleaned_body"`
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
	msg := profanityFilter(params.Body)	
	respondWithJSON(w, http.StatusOK, validResponse{
		CleanedBody: msg,
	})
}

func profanityFilter(msg string) string {	
	profanities := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}
	
	words := strings.Fields(msg)
	for i, word := range words {
		if slices.Contains(profanities, strings.ToLower(word)){
			words[i] = "****"
		}
	}
	msg = strings.Join(words, " ")
	return msg
}