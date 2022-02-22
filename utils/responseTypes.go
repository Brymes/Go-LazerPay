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
