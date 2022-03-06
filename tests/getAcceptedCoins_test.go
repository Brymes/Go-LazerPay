package tests

import (
	"github.com/Brymes/Go-LazerPay/services"
	"testing"
)

func TestAcceptedCoins(t *testing.T) {
	testCase := services.GetAcceptedCoins{}

	res, err := testCase.GetCoins(testKeys)
	if err != nil {
		t.Errorf("Test: Get All Coins failed: \n Error was: %v \n Data Returned is: %v", err, res)
	}
}
