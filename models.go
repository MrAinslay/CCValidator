package main

type BinData struct {
	Result  int    `json:"result"`
	Message string `json:"message"`
	Data    struct {
		Card struct {
			Scheme    string `json:"scheme"`
			Type      string `json:"type"`
			Category  string `json:"category"`
			Length    int    `json:"length"`
			Checkluhn int    `json:"checkluhn"`
			Cvvlen    int    `json:"cvvlen"`
		} `json:"card"`
		Country struct {
			Name         string `json:"name"`
			Code         string `json:"code"`
			Flag         string `json:"flag"`
			Currency     string `json:"currency"`
			CurrencyCode string `json:"currency_code"`
		} `json:"country"`
		Bank struct {
			Name    string `json:"name"`
			Website string `json:"website"`
			Phone   string `json:"phone"`
		} `json:"bank"`
	} `json:"data"`
}

type BankInfo struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}

type CardInfo struct {
	Scheme   string `json:"scheme"`
	Type     string `json:"type"`
	Category string `json:"category"`
}

type CountryInfo struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type DataInfo struct {
	CardInfo struct {
		Scheme   string `json:"scheme"`
		Type     string `json:"type"`
		Category string `json:"category"`
	} `json:"card_info"`
	Bank struct {
		Name    string `json:"name"`
		Website string `json:"website"`
		Phone   string `json:"phone"`
	} `json:"bank"`
	Country struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"country"`
}

type BinResponseData struct {
	Result string `json:"result"`
	Data   struct {
		CardInfo struct {
			Scheme   string `json:"scheme"`
			Type     string `json:"type"`
			Category string `json:"category"`
		} `json:"card_info"`
		Bank struct {
			Name    string `json:"name"`
			Website string `json:"website"`
			Phone   string `json:"phone"`
		} `json:"bank"`
		Country struct {
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"country"`
	} `json:"data"`
}
