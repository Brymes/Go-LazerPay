package tests

import (
	"LazerPay-Go-SDK/services"
	"LazerPay-Go-SDK/utils"
	"testing"
)

func TestAcceptedCoins(t *testing.T) {
	keys := utils.ApiKeys{
		PubKey: "pk_live_0N24k7lsrr7NGfrDQpIjPGy9z61LkXjUqxX3r99XblXHemwMht",
		SecKey: "sk_alals",
	}

	tmp := services.GetAcceptedCoins{}
	res, err := tmp.Get(keys)
	if err != nil {
		t.Errorf("Get All Coins failed: \n Error was: %v \n Data Returned is: %v", err, res)
	}
}
