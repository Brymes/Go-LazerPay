package tests

import (
	"LazerPay-Go-SDK/services"
	"LazerPay-Go-SDK/utils"
	"testing"
)

func TestPayoutSuccess(t *testing.T) {
	keys := utils.ApiKeys{
		PubKey: "pk_live_0N24k7lsrr7NGfrDQpIjPGy9z61LkXjUqxX3r99XblXHemwMht",
		SecKey: "pk_test_QVaSyqlCqhpY585oHV3hc01nfMZfOzhy5YpPIOc0LywoWK4EVt",
	}
	req := services.TransferFunds{
		Amount:     1,
		Recipient:  "0xF378c952d5266eF8e1783521a1395Fe40cDCe55B",
		Coin:       "USDT",
		Blockchain: "Binance Smart Chain",
	}

	res, err := req.Transfer(keys)
	if err != nil {
		t.Errorf("TransferFunds failed failed: \n Error was: %v \n Data Returned is: %v", err, res)
	}
}

func TestPayoutFailure(t *testing.T) {}
