package services

import (
	u "LazerPay-Go-SDK/utils"
	"bytes"
	"encoding/json"
	"net/http"
)

type InitializeTransaction struct {
	CustomerName  string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
	Coin          string `json:"coin"`
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
}

func (info *InitializeTransaction) Initialize(keys u.ApiKeys) (u.IniitalizeTransactionResponse, error) {
	var data u.IniitalizeTransactionResponse

	url := u.ApiUrlTransferFunds
	payload, err := json.Marshal(&info)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	resp := u.MakeReq(*req, keys, false)

	err = json.Unmarshal(resp, &data)
	if err != nil {
		return data, err
	}

	return data, err
}
