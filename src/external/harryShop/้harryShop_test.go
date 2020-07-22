package harryShop

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/oommi04/shibabookbackend/src/drivers/fastHttpDriver/mocks"
)

type HarryShopServiceSuite struct {
	suite.Suite

	fastHttp           *mocks.FastHttpClient
	service            *HarryShopClient
	integrationService *HarryShopClient
}

func Test_Google_Service_Suite(t *testing.T) {
	suite.Run(t, new(HarryShopServiceSuite))
}

func (suite *HarryShopServiceSuite) SetupTest() {
	suite.fastHttp = &mocks.FastHttpClient{}
	suite.service = New("localhost",  5).setHttpClient(suite.fastHttp)
	suite.integrationService = New("https://api.jsonbin.io/b",  5)
}
