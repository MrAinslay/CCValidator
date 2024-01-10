package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (cfg *apiConfig) getBinData(bin string) (BinResponseData, error) {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	url := cfg.baseUrl + bin + "?api_key=" + cfg.apiKey
	rsp, err := httpClient.Get(url)
	if err != nil {
		return BinResponseData{}, err
	}

	decoder := json.NewDecoder(rsp.Body)
	dat := BinData{}
	if err := decoder.Decode(&dat); err != nil {
		return BinResponseData{}, err
	}

	return BinResponseData{
		Result: "Valid credit card number",
		Data: DataInfo{
			CardInfo: CardInfo{
				Scheme:   dat.Data.Card.Scheme,
				Type:     dat.Data.Card.Type,
				Category: dat.Data.Card.Category,
			},
			Bank: BankInfo{
				Name:    dat.Data.Bank.Name,
				Website: dat.Data.Bank.Website,
				Phone:   dat.Data.Bank.Phone,
			},
			Country: CountryInfo{
				Name: dat.Data.Country.Name,
				Code: dat.Data.Country.Code,
			},
		},
	}, nil
}
