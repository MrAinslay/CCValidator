package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerAccountRange(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		AccountRng string `json:"account_range"`
	}

	decoder := json.NewDecoder(r.Body)
	prm := parameters{}
	if err := decoder.Decode(&prm); err != nil {
		log.Printf("Error decoding request: %v", err)
		respondWithErr(w, 500, "Couldn't decode request")
		return
	}

	if prm.AccountRng == "" {
		log.Println("Invalid json body")
		respondWithErr(w, 500, "Invalid json format")
		return
	}

	dat, err := cfg.getBinData(prm.AccountRng)
	if err != nil {
		log.Printf("Error retrieving BIN data from api: %v", err)
		respondWithErr(w, 404, "Couldn't retrieve BIN data")
		return
	}
	dat.Result = "Valid account range"

	respondWithJSON(w, 200, dat)
}
