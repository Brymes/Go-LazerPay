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
