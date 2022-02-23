package utils

import "time"

type AcceptedCoins struct {
	Message string `json:"message"`
	Data    []struct {
		Id         string    `json:"id"`
		Name       string    `json:"name"`
		Symbol     string    `json:"symbol"`
		Logo       string    `json:"logo"`
		Address    string    `json:"address"`
		Network    string    `json:"network"`
		Blockchain string    `json:"blockchain"`
		Status     string    `json:"status"`
		CreatedAt  time.Time `json:"createdAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
	} `json:"data"`
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
}

type TransferResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  string `json:"status"`
	Data    struct {
		Id              string    `json:"id"`
		CreatedAt       time.Time `json:"createdAt"`
		UpdatedAt       time.Time `json:"updatedAt"`
		TransactionHash string    `json:"transactionHash"`
		WalletAddress   string    `json:"walletAddress"`
		Coin            string    `json:"coin"`
		Amount          int       `json:"amount"`
	} `json:"data"`
	StatusCode int `json:"statusCode"`
}

type VerifyTransactionResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       struct {
		Id               string `json:"id"`
		Reference        string `json:"reference"`
		SenderAddress    string `json:"senderAddress"`
		RecipientAddress string `json:"recipientAddress"`
		ActualAmount     int    `json:"actualAmount"`
		AmountPaid       int    `json:"amountPaid"`
		FiatAmount       int    `json:"fiatAmount"`
		Coin             string `json:"coin"`
		Currency         string `json:"currency"`
		Hash             string `json:"hash"`
		BlockNumber      int    `json:"blockNumber"`
		Type             string `json:"type"`
		Status           string `json:"status"`
		Network          string `json:"network"`
		Blockchain       string `json:"blockchain"`
		Customer         struct {
			Id            string `json:"id"`
			CustomerName  string `json:"customerName"`
			CustomerEmail string `json:"customerEmail"`
			CustomerPhone int    `json:"customerPhone"`
			Network       string `json:"network"`
		} `json:"customer"`
	} `json:"data"`
}

type IniitalizeTransactionResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		BusinessName  string `json:"businessName"`
		BusinessEmail string `json:"businessEmail"`
		BusinessLogo  string `json:"businessLogo"`
		CustomerName  string `json:"customerName"`
		CustomerEmail string `json:"customerEmail"`
		Address       string `json:"address"`
		Coin          string `json:"coin"`
		CryptoAmount  int    `json:"cryptoAmount"`
		Currency      string `json:"currency"`
		FiatAmount    int    `json:"fiatAmount"`
		Network       string `json:"network"`
	} `json:"data"`
	StatusCode int `json:"statusCode"`
}