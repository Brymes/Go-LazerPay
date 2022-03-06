package tests

import (
	"github.com/Brymes/Go-LazerPay/utils"
	"os"
)

var testKeys = utils.ApiKeys{
	PubKey: os.Getenv("LAZERPAY_PUB_KEY"),
	SecKey: os.Getenv("LAZERPAY_SEC_KEY"),
}
