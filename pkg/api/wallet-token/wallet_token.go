package wallet_token

import (
	"fmt"
	"github.com/cryptounifier/go-sdk/pkg/api/base"
	"github.com/cryptounifier/go-sdk/pkg/models"
)

type WalletTokenAPIClient interface {
	base.BaseAPIClient
	//SendTokenTransaction creates and broadcasts a token transaction
	SendTokenTransaction(destinations map[string]interface{}, feePerByte float64, extraField string) (models.Response, error)
	//EstimateTokenFee estimates a final token transaction fee and it fee per byte cost
	EstimateTokenFee(destinations map[string]interface{}, feePerByte float64, extraField string) (models.Response, error)
	//GetTokenBalance gets confirmed and unconfirmed cryptocurrency token balance
	GetTokenBalance() (models.Response, error)
}
type client struct {
	walletKey    string
	secretKey    string
	cryptoSymbol string
	tokenSymbol  string
	headers      map[string]string
	suffix       string
	*base.BaseAPI
}

func NewWalletTokenApiClientFactory(walletKey, secretKey, cryptoSymbol, tokenSymbol string) WalletTokenAPIClient {
	c := client{
		walletKey:    walletKey,
		secretKey:    secretKey,
		cryptoSymbol: cryptoSymbol,
		suffix:       fmt.Sprintf("wallet/%s/token/%s", cryptoSymbol, tokenSymbol),
	}
	c.headers = map[string]string{
		"X-Wallet-Key": walletKey,
		"X-Secret-Key": secretKey,
	}

	c.BaseAPI = base.NewBaseApi(c.suffix, c.headers)

	return c
}

func (c client) SendTokenTransaction(destinations map[string]interface{}, feePerByte float64, extraField string) (models.Response, error) {
	body := map[string]interface{}{
		"destinations": destinations,
		"fee_per_byte": feePerByte,
		"extra_field":  extraField,
	}

	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.POST, "send-transaction", body)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) EstimateTokenFee(destinations map[string]interface{}, feePerByte float64, extraField string) (models.Response, error) {
	body := map[string]interface{}{
		"destinations": destinations,
		"fee_per_byte": feePerByte,
		"extra_field":  extraField,
	}

	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.POST, "estimate-fee", body)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) GetTokenBalance() (models.Response, error) {
	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.GET, "balance", nil)
	if err != nil{
		return models.Response{}, err
	}
	return *resp, nil
}
