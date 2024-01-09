package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/MrAinslay/CCValidator/internal/validator"
	"github.com/mastercard/oauth1-signer-go/interceptor"
)

func (cfg *apiConfig) handlerValidateCC(w http.ResponseWriter, r *http.Request) {
	const baseURL = "https://sandbox.api.mastercard.com"
	type parameters struct {
		CCNumber int `json:"credit_card_number"`
	}
	type rsp struct {
		Body       string `json:"message"`
		AuthHeader string `json:"auth_header"`
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
	w.Header().Add("Content-Type", "application/json")
	httpClient, _ := interceptor.GetHttpClient(cfg.consumerKey, cfg.pathToCerts, cfg.keyPass)
	jsonBody := []byte(`{"accountRange": "53514204"}`)
	bodyReader := bytes.NewReader(jsonBody)
	resp, err := httpClient.Post(baseURL+"/bin-ranges/account-searches", "application/json", bodyReader)
	if err != nil {
		log.Printf("Error posting to api: %v", err)
		respondWithErr(w, 500, "Couldn't post to api")
		return
	}
	log.Println(&resp.Body, *resp, resp.Body, resp)
	dcder := json.NewDecoder(resp.Body)
	rspns := BINInfo{}
	if err := dcder.Decode(&rspns); err != nil {
		log.Printf("Error decoding response: %v", err)
		respondWithErr(w, 500, "Couldn't decode response")
		return
	}

	respondWithJSON(w, 200, rspns)
}
