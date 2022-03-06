package LazerPay

import (
	"github.com/Brymes/Go-LazerPay/services"
	"github.com/Brymes/Go-LazerPay/utils"
)

type LazerPay struct {
	Keys utils.ApiKeys
}

func (l LazerPay) getAcceptedCoins() (utils.AcceptedCoins, error) {
	return services.GetAcceptedCoins{}.GetCoins(l.Keys), nil
}

func (l LazerPay) initTransaction(transactionPayload services.InitializeTransaction) (utils.IniitalizeTransactionResponse, error) {
	return transactionPayload.InitTrans(l.Keys), nil
}

func (l LazerPay) transferFunds(transferPayload services.TransferFunds) (utils.TransferResponse, error) {
	return transferPayload.Transfer(l.Keys), nil
}

func (l LazerPay) confirmPayment(confirmPayload services.ConfirmTransaction) (utils.VerifyTransactionResponse, error) {
	return confirmPayload.Confirm(l.Keys), nil
}