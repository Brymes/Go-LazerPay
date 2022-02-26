package tests

import (
	"LazerPay-Go-SDK/services"
	"testing"
)

func TestInitTransaction(t *testing.T) {
	testCase := services.InitializeTransaction{
		CustomerName:  "Abdulfatai Suleiman",
		CustomerEmail: "staticdev20046@gmail.com",
		Coin:          "DAI",
		Currency:      "USD",
		Amount:        "1000",
	}

	res, err := testCase.InitTrans(testKeys)

	if err != nil {
		t.Errorf("Test: Initialize Transactions Failed \n Error was: %v \n Data Returned is: %v", err, res)
	}
}
