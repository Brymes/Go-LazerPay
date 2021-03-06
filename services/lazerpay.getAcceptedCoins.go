package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	u "github.com/Brymes/Go-LazerPay/utils"
	"net/http"
)

type GetAcceptedCoins struct{}

func (gcs GetAcceptedCoins) GetCoins(keys u.ApiKeys) (u.AcceptedCoins, error) {
	var data u.AcceptedCoins
	var b bytes.Buffer

	url := u.ApiUrlGetAcceptedCoins

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
