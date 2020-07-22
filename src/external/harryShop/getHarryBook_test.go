package harryShop

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

func (suite *HarryShopServiceSuite) TestHarryShopClient_GetHarryBook_Success() {
	req, _ := suite.service.buildGetRequest(pathHarryBook)
	resp := &fasthttp.Response{}

	suite.fastHttp.On("DoTimeout", req, resp, 5*time.Second).Once().Run(func(args mock.Arguments) {
		resp := args[1].(*fasthttp.Response)
		resp.SetStatusCode(http.StatusOK)

		respBody := HarryBookList{Books: []*Books{
				{
					ID: "9781408855652",
					Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855652.jpg",
					Price: "350",
					Title: "Harry Potter and the Philosopher's Stone (I)",
				},
		}}
		creatorJSON, _ := json.Marshal(respBody)
		resp.SetBody(creatorJSON)
	}).Return(nil)

	_, err := suite.service.GetHarryBook()

	suite.NoError(err)
	suite.fastHttp.AssertExpectations(suite.T())
}

func (suite *HarryShopServiceSuite) TestHarryShopClient_GetHarryBook_Error() {
	req, _ := suite.service.buildGetRequest(pathHarryBook)
	resp := &fasthttp.Response{}

	suite.fastHttp.On("DoTimeout", req, resp, 5*time.Second).Once().Run(func(args mock.Arguments) {
		resp := args[1].(*fasthttp.Response)
		resp.SetStatusCode(http.StatusBadRequest)
	}).Return(nil)

	_, err := suite.service.GetHarryBook()

	suite.Error(err)
	suite.Equal(ErrorUnableRequestGetHarryBook, err)
	suite.fastHttp.AssertExpectations(suite.T())
}
