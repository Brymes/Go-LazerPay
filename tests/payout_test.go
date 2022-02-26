package tests

import (
	"LazerPay-Go-SDK/services"
	"testing"
)

func TestPayoutSuccess(t *testing.T) {
	req := services.TransferFunds{
		Amount:     1,
		Recipient:  "0xF378c952d5266eF8e1783521a1395Fe40cDCe55B",
		Coin:       "USDT",
		Blockchain: "Binance Smart Chain",
	}

	res, err := req.Transfer(testKeys)
	if err != nil {
		t.Errorf("TransferFunds failed failed: \n Error was: %v \n Data Returned is: %v", err, res)
	}
}

func TestPayoutFailure(t *testing.T) {}
