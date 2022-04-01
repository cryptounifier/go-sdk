# CryptoUnifier GO SDK

A simple GO SDK for interacting with [Crypto Unifier](https://cryptounifier.io) API V1.

## Installation

You can install the package from github:

```bash
go get github.com/cryptounifier/go-sdk
```

## Usage

### Using the Wallet API client

You can use the `WalletAPI` class for convenient access to API methods. Some are defined in the code:

```go
import (
	"github.com/cryptounifier/go-sdk"
    "log"
)

client := cryptounifier.WalletAPI("WALLET_KEY", "SECRET_KEY", "btc")

resp, err := client.GetBalance()
if err != nil {
	log.Fatal(err)
}

log.Println("message: ", resp.Message)
```

### Using the Merchant API client

You can use the `MerchantAPI` class for convenient access to API methods. Some are defined in the code:

```go
import (
	"github.com/cryptounifier/go-sdk"
    "log"
)

client := cryptounifier.MerchantAPI("MERCHANT_KEY", "SECRET_KEY")

resp, err := client.CreateInvoice([...]string{"btc", "bch", "eth"})
if err != nil {
	log.Fatal(err)
}

log.Println("message:", resp.Message)
```

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.