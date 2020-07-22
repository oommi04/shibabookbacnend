package harryShop

import (
	"errors"
	"github.com/tkhamsila/shibabookbackend/src/domains/productDomain"
	"github.com/tkhamsila/shibabookbackend/src/drivers/fastHttpDriver"
	"github.com/valyala/fasthttp"
)

var (
	ErrorUnableCreateRequest                    = errors.New("unable create request from path")
	ErrorUnableRequestGetHarryBook          = errors.New("unable request get harry book")
)

type HarryShopClientInterface interface {
	GetHarryBook() ([]*productDomain.Product, error)
}

type HarryShopClient struct {
	httpClient fastHttpDriver.FastHttpClient

	endpoint string
	timeout  int
}

func New(endpoint string, timeout int) *HarryShopClient {
	return &HarryShopClient{
		endpoint:   endpoint,
		httpClient: &fasthttp.Client{},
		timeout:    timeout,
	}
}

func (client *HarryShopClient) setHttpClient(httpClient fastHttpDriver.FastHttpClient) *HarryShopClient {
	client.httpClient = httpClient

	return client
}

func (client *HarryShopClient) buildGetRequest(path string) (*fasthttp.Request, error) {
	req := &fasthttp.Request{}

	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")

	fullEndpoint := client.endpoint + path

	req.SetRequestURI(fullEndpoint)

	return req, nil
}