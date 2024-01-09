package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MrAinslay/CCValidator/internal/validator"
)

func (cfg *apiConfig) handlerValidateCC(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		CCNumber int `json:"credit_card_number"`
	}
	type rsp struct {
		Body string `json:"message"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		log.Printf("Error decoding request body: %v", err)
		respondWithErr(w, 500, "Couldn't decode response body")
		return
	}

	if params.CCNumber == 0 {
		respondWithErr(w, 500, "Invalid json body")
		return
	}

	ok := validator.ValidateNum(params.CCNumber)
	if !ok {
		respondWithErr(w, 500, "Invalid credit card number")
		return
	}

	respondWithJSON(w, 200, rsp{Body: "Valid credit card number"})
}
