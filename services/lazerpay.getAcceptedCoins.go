package services

import (
	u "LazerPay-Go-SDK/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GetAcceptedCoins struct {
	Url string
}

func (gcs *GetAcceptedCoins) Get(keys u.ApiKeys) (u.AcceptedCoins, error) {
	var data u.AcceptedCoins
	var b bytes.Buffer

	req, err := http.NewRequest("GET", gcs.Url, &b)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	resp := u.MakeReq(*req, keys)

	err = json.Unmarshal(resp, &data)
	if err != nil {
		return data, err
	}

	return data, err
}
