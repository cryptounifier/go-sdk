package base

import (
	"bytes"
	"encoding/json"
	"github.com/cryptounifier/go-sdk/pkg/config"
	"github.com/cryptounifier/go-sdk/pkg/models"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var RequestMethod = struct {
	GET  string
	POST string
}{
	"GET",
	"POST",
}

type BaseAPIClient interface {
	SetApiURL(apiURL string)
}

type BaseAPI struct {
	defaultURL string
	apiKey     string
	secretKey  string
	suffix     string
	headers    map[string]string
	baseURI    string
	client     http.Client
}

func NewBaseApi(suffix string, headers map[string]string) *BaseAPI {
	b := &BaseAPI{
		defaultURL: config.DefaultURL,
		suffix:     suffix,
		headers:    headers,
	}
	b.baseURI = strings.Join([]string{b.defaultURL, b.suffix}, "/")
	b.client = http.Client{
		Timeout: 10 * time.Second,
	}

	return b
}

func (b *BaseAPI) SetApiURL(apiURL string) {
	b.baseURI = strings.Join([]string{apiURL, b.suffix}, "/")
}

func (b *BaseAPI) ExecuteRequest(method, uri string, payload interface{}) (*models.Response, error) {
	url := b.baseURI + "/" + uri

	payloadInBytes, err := json.Marshal(payload)
	if err !=nil{
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadInBytes))
	if err != nil{
		return nil, err
	}

	for key, val := range b.headers{
		req.Header.Add(key, val)
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return decodeResponseBody(body)
}

func decodeResponseBody(body []byte) (*models.Response, error){
	var resp models.Response
	err := json.Unmarshal(body, &resp)
	if err != nil{
		return nil, err
	}
	return &resp, nil
}