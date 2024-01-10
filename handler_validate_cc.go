package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MrAinslay/CCValidator/internal/validator"
)

func (cfg *apiConfig) handlerValidateCC(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		CCNumber string `json:"credit_card_number"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		log.Printf("Error decoding request body: %v", err)
		respondWithErr(w, 500, "Couldn't decode response body")
		return
	}

	if params.CCNumber == "" {
		respondWithErr(w, 500, "Invalid json body")
		return
	}

	ok := validator.ValidateNum(params.CCNumber)
	if !ok {
		respondWithErr(w, 500, "Invalid credit card number")
		return
	}

	strNum := fmt.Sprint(params.CCNumber)
	digits := strings.Split(strNum, "")
	accountRange := ""
	for i := 0; i < 6; i++ {
		accountRange += digits[i]
	}
	dat, err := cfg.getBinData(accountRange)
	if err != nil {
		log.Printf("Error getting bin data: %v", err)
		respondWithErr(w, 500, "Couldn't get BIN data")
		return
	}

	respondWithJSON(w, 200, dat)
}
