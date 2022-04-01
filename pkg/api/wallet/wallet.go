package wallet

import (
	"fmt"
	"github.com/cryptounifier/go-sdk/pkg/api/base"
	"github.com/cryptounifier/go-sdk/pkg/models"
)

type WalletAPIClient interface {
	base.BaseAPIClient
	//SendTransaction creates and broadcasts a transaction
	SendTransaction(destinations map[string]interface{}, feePerByte float64, extraField string) (models.Response, error)
	//EstimateFee estimates a final transaction fee and it fee per byte cost
	EstimateFee(destinations map[string]interface{}, feePerByte float64, extraField string) (models.Response, error)
	//ValidateAddress check if an address list is currently valid for selected cryptocurrency
	ValidateAddress(addresses []interface{}) (models.Response, error)
	//GetBalance get confirmed and unconfirmed cryptocurrency balance
	GetBalance() (models.Response, error)
	//GetDepositAddresses get the list of cryptocurrency deposit addresses
	GetDepositAddresses() (models.Response, error)
	//GetBlockchainInfo get information about the current state of the cryptocurrency blockchain and the sync percentage
	//from the connected node
	GetBlockchainInfo() (models.Response, error)
	//GetTransactionInfo
	GetTransactionInfo(txID string) (models.Response, error)
}

type client struct {
	walletKey    string
	secretKey    string
	cryptoSymbol string
	headers      map[string]string
	suffix       string
	*base.BaseAPI
}

func NewWalletApiClientFactory(walletKey, secretKey, cryptoSymbol string) WalletAPIClient {
	c := client{
		walletKey:    walletKey,
		secretKey:    secretKey,
		cryptoSymbol: cryptoSymbol,
		suffix:       fmt.Sprintf("wallet/%s", cryptoSymbol),
	}
	c.headers = map[string]string{
		"X-Wallet-Key": walletKey,
		"X-Secret-Key": secretKey,
	}

	c.BaseAPI = base.NewBaseApi(c.suffix, c.headers)

	return c
}

func (c client) SendTransaction(destinations map[string]interface{}, feePerByte float64, extraField string) (models.Response, error) {
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

func (c client) EstimateFee(destinations map[string]interface{}, feePerByte float64, extraField string) (models.Response, error) {
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

func (c client) ValidateAddress(addresses []interface{}) (models.Response, error) {
	body := map[string]interface{}{
		"addresses": addresses,
	}

	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.POST, "validate-addresses", body)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) GetBalance() (models.Response, error) {
	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.GET, "balance", nil)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) GetDepositAddresses() (models.Response, error) {
	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.GET, "deposit-addresses", nil)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) GetBlockchainInfo() (models.Response, error) {
	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.GET, "blockchain-info", nil)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) GetTransactionInfo(txID string) (models.Response, error) {
	body := map[string]interface{}{
		"txid": txID,
	}
	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.GET, "transaction-info", body)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}
