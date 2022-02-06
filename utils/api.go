package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeReq(r http.Request, keys ApiKeys) []byte {

	// Use Struct to pass in API Keys
	if keys.SecKey != "" {
		token := fmt.Sprintf(`Bearer %v`, keys.SecKey)
		r.Header.Add("Authorization", token)
	}

	r.Header.Add("x-api-key", keys.PubKey)

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
	return resp
}
