package utils

import "fmt"

const API_URL = "https://api.lazerpay.engineering/api/v1"

var API_URL_INIT_TRANSACTION = fmt.Sprintf("%v/transaction/initialize", API_URL)

var API_URL_CONFIRM_TRANSACTION = fmt.Sprintf("%v/transaction/verify", API_URL)

var API_URL_GET_ACCEPTED_COINS = fmt.Sprintf("%v/coins", API_URL)
var API_URL_TRANSFER_FUNDS = fmt.Sprintf("%v/transfer", API_URL)
