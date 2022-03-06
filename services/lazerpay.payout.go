package services

import (
	u "github.com/Brymes/Go-LazerPay/utils"
	"bytes"
	"encoding/json"
	"net/http"
)

type TransferFunds struct {
	Amount     int    `json:"amount"`
	Recipient  string `json:"recipient"`
	Coin       string `json:"coin"`
	Blockchain string `json:"blockchain"`
}

func (tf *TransferFunds) Transfer(keys u.ApiKeys) (u.TransferResponse, error) {
	var data u.TransferResponse

	url := u.ApiUrlTransferFunds
	payload, err := json.Marshal(&tf)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	resp := u.MakeReq(*req, keys, true)

	err = json.Unmarshal(resp, &data)
	if err != nil {
		return data, err
	}

	return data, err
}
