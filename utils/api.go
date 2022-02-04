package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeReq(r http.Request, apiPubKey, apiSecKey string) {
	if apiSecKey != "" {
		token := fmt.Sprintf(`Bearer %v`, apiSecKey)
		r.Header.Add("Authorization", token)
	}

	r.Header.Add("x-api-key", apiPubKey)

	client := &http.Client{}

	res, err := client.Do(&r)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(resp)
}
