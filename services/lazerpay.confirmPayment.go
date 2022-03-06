package services

import (
	u "github.com/Brymes/Go-LazerPay/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ConfirmTransaction struct {
	Identifier string
}

func (id *ConfirmTransaction) Confirm(keys u.ApiKeys) (u.VerifyTransactionResponse, error) {
	var data u.VerifyTransactionResponse
	var b bytes.Buffer

	url := fmt.Sprintf("%s/%s", u.ApiUrlGetAcceptedCoins, id.Identifier)

	req, err := http.NewRequest("GET", url, &b)

	if err != nil {
		fmt.Println(err)
		return data, err
	}
	resp := u.MakeReq(*req, keys, false)

	err = json.Unmarshal(resp, &data)
	if err != nil {
		return data, err
	}

	return data, err
}
