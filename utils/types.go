package utils

type TransactionPayloadData struct {
	Amount        string
	Coin          string
	Currency      string
	CustomerEmail string
	CustomerName  string
}

type ConfirmTransactionPayloadData struct {
	address string
}

type TransferFundsPayloadData struct {
	Amount     string
	Recipient  string
	Coin       string
	Blockchain string
}

type ApiKeys struct {
	PubKey, SecKey string
}
