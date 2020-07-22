package harryShop

func (suite *HarryShopServiceSuite) TestHarryShopClient_Integration_GetHarryBook_Success() {
	_, err := suite.integrationService.GetHarryBook()
	suite.NoError(err)
	suite.fastHttp.AssertExpectations(suite.T())
}
