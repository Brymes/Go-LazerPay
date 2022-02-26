package tests

import (
	"LazerPay-Go-SDK/services"
	"testing"
)

func TestConfirmPayment(t *testing.T) {
	testCase := services.ConfirmTransaction{
		Identifier: "0x80179d818235E87aFfFdb6FD89C97452F3a22746",
	}

	res, err := testCase.Confirm(testKeys)

	if err != nil {
		t.Errorf("Test: Confirm Transaction Failed \n Error was: %v \n Data Returned is: %v", err, res)
	}
}
