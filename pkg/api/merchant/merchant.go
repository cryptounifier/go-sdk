package merchant

import (
	"fmt"
	"github.com/cryptounifier/go-sdk/pkg/api/base"
	"github.com/cryptounifier/go-sdk/pkg/models"
)

//MerchantAPIClient
type MerchantAPIClient interface {
	base.BaseAPIClient
	//CreateInvoice creates an invoice to charge for a product or service
	CreateInvoice(cryptocurrencies []string, currency string, targetValue float64, title string, description string) (models.Response, error)
	//EstimateInvoicePrice estimates invoice price value for multiple cryptocurrencies
	EstimateInvoicePrice(cryptocurrencies []string, currency string, targetValue float64) (models.Response, error)
	//ProcessInvoices manually processes expired invoices in order to update received amount
	ProcessInvoices(invoiceHashes []interface{}) (models.Response, error)
	//ForwardInvoices manually forwards invoice funds
	ForwardInvoices(invoiceHashes []interface{}) (models.Response, error)
	//GenerateInvoiceAddress generates invoice address for a specific cryptocurrency
	GenerateInvoiceAddress(invoiceHash string, cryptocurrency string) (models.Response, error)
	//RecoverInvoicePrivateKey recovers your private key for a specific cryptocurrency
	RecoverInvoicePrivateKey(invoiceHash string, cryptocurrency string) (models.Response, error)
	//InvoiceInfo gets invoice information and its current status
	InvoiceInfo(invoiceHash string) (models.Response, error)
}

type client struct {
	merchantKey string
	secretKey   string
	headers     map[string]string
	suffix      string
    *base.BaseAPI
}

func NewMerchantApiClientFactory(merchantKey, secretKey string) MerchantAPIClient {
	c := client{
		merchantKey: merchantKey,
		secretKey:   secretKey,
		suffix:      "merchant",
	}
	c.headers = map[string]string{
		"X-Merchant-Key": merchantKey,
		"X-Secret-Key":   secretKey,
	}
	c.BaseAPI = base.NewBaseApi(c.suffix, c.headers)

	return c
}

func (c client) CreateInvoice(cryptocurrencies []string, currency string, targetValue float64, title string, description string) (models.Response, error) {
	body := map[string]interface{}{
		"cryptocurrencies": cryptocurrencies,
		"currency":         currency,
		"target_value":     targetValue,
		"title":            title,
		"description":      description,
	}

	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.POST, "create-invoice", body)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) EstimateInvoicePrice(cryptocurrencies []string, currency string, targetValue float64) (models.Response, error) {
	body := map[string]interface{}{
		"cryptocurrencies": cryptocurrencies,
		"currency":         currency,
		"target_value":     targetValue,
	}
	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.POST, "estimate-invoice-price", body)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) ProcessInvoices(invoiceHashes []interface{}) (models.Response, error) {
	body := map[string]interface{}{
		"invoice_hashes": invoiceHashes,
	}
	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.POST, "process-invoices", body)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) ForwardInvoices(invoiceHashes []interface{}) (models.Response, error) {
	body := map[string]interface{}{
		"invoice_hashes": invoiceHashes,
	}
	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.POST, "forward-invoices", body)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) GenerateInvoiceAddress(invoiceHash string, cryptocurrency string) (models.Response, error) {
	body := map[string]interface{}{
		"invoice_hash":   invoiceHash,
		"cryptocurrency": cryptocurrency,
	}
	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.POST, "generate-invoice-address", body)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) RecoverInvoicePrivateKey(invoiceHash string, cryptocurrency string) (models.Response, error) {
	body := map[string]interface{}{
		"invoice_hash":   invoiceHash,
		"cryptocurrency": cryptocurrency,
	}
	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.POST, "recover-invoice-private-key", body)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}

func (c client) InvoiceInfo(invoiceHash string) (models.Response, error) {
	uri := fmt.Sprintf("invoice-info?invoice_hash=%s", invoiceHash)
	resp, err := c.BaseAPI.ExecuteRequest(base.RequestMethod.GET, uri, nil)
	if err != nil {
		return models.Response{}, err
	}
	return *resp, nil
}
