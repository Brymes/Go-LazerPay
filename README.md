# Go-LazerPay

Unofficial Golang SDK for LazerPay

## Installation

`go get github.com/Brymes/Go-LazerPay`

## Usage

```
import (
    "os"
    "github.com/Brymes/Go-LazerPay.git"
)

var keys = LazerPay.utils.ApiKeys{
    PubKey: os.Getenv("LAZERPAY_PUB_KEY"),
 SecKey: os.Getenv("LAZERPAY_SEC_KEY"),    
}

var client = LazerPay{ keys }
```

### Get Accepted Coins

```
allCoins := client.getAcceptedCoins()
```

### Confirm Transaction

```
txid := LazerPay.services.ConfirmTransaction{
    Identifier: "<Your_Transaction_ID>"
}

confirm := client.getAcceptedCoins(txid)
```

## TESTING

```.
$ export LAZERPAY_SEC_KEY="<key_here>" 
$ export LAZERPAY_PUB_KEY="<key_here>" 

$ cd tests && go test -v
=== RUN   TestAcceptedCoins
--- PASS: TestAcceptedCoins (4.82s)
PASS
ok      github.com/Brymes/Go-LazerPay/tests   7.891s
```
