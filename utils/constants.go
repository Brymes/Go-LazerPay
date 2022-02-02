package utils

import "fmt"

const ApiUrl = "https://api.lazerpay.engineering/api/v1"

var ApiUrlInitTransaction = fmt.Sprintf("%v/transaction/initialize", ApiUrl)

var ApiUrlConfirmTransaction = fmt.Sprintf("%v/transaction/verify", ApiUrl)

var ApiUrlGetAcceptedCoins = fmt.Sprintf("%v/coins", ApiUrl)

var ApiUrlTransferFunds = fmt.Sprintf("%v/transfer", ApiUrl)
