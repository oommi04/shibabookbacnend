package harryShop

import (
	"encoding/json"
	"github.com/oommi04/shibabookbackend/src/domains/productDomain"
	"github.com/oommi04/shibabookbackend/src/utils/common"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

type Books struct {
	Cover string `json:"cover"`
	Price string `json:"price"`
	Title string `json:"title"`
	ID    string `json:"id"`
}

type HarryBookList struct {
	Books []*Books `json:"books"`
}

var pathHarryBook = "/5e69b564d2622e7011565547"

func (client *HarryShopClient) GetHarryBook() ([]*productDomain.Product, error) {
	req, err := client.buildGetRequest(pathHarryBook)

	if err != nil {
		return nil, ErrorUnableCreateRequest
	}

	resp := &fasthttp.Response{}

	err = client.httpClient.DoTimeout(req, resp, time.Duration(client.timeout)*time.Second)

	if err != nil {
		return nil, ErrorUnableRequestGetHarryBook
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, ErrorUnableRequestGetHarryBook
	}

	respBody := HarryBookList{}
	_ = json.Unmarshal(resp.Body(), &respBody)

	productEntity := []*productDomain.Product{}

	for _,item := range respBody.Books {
		productEntity = append(productEntity, &productDomain.Product{
			Price: float32(common.StringToInt(item.Price)),
			Image: item.Cover,
			Name: item.Title,
			From: "HarryShop",
		})
	}

	return productEntity, nil
}