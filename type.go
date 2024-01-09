package main

type BINInfo struct {
	LowAccountRange  int64  `json:"lowAccountRange"`
	HighAccountRange int64  `json:"highAccountRange"`
	BinNum           string `json:"binNum"`
	BinLength        string `json:"binLength"`
	AcceptanceBrand  string `json:"acceptanceBrand"`
	Ica              string `json:"ica"`
	CustomerName     string `json:"customerName"`
	Country          struct {
		Code   int    `json:"code"`
		Alpha3 string `json:"alpha3"`
		Name   string `json:"name"`
	} `json:"country"`
	LocalUse                    string `json:"localUse"`
	ProductCode                 string `json:"productCode"`
	ProductDescription          string `json:"productDescription"`
	GovernmentRange             string `json:"governmentRange"`
	NonReloadableIndicator      string `json:"nonReloadableIndicator"`
	AnonymousPrepaidIndicator   string `json:"anonymousPrepaidIndicator"`
	CardholderCurrencyIndicator string `json:"cardholderCurrencyIndicator"`
	FundingSource               string `json:"fundingSource"`
	ConsumerType                string `json:"consumerType"`
	Affiliate                   string `json:"affiliate"`
}
