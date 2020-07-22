package app

import (
	"context"
	"github.com/oommi04/shibabookbackend/src/domains/productDomain"
	"github.com/oommi04/shibabookbackend/src/external/harryShop"
	"github.com/oommi04/shibabookbackend/src/usecase/productUsecase"
)

func SetUpHarryShop() harryShop.HarryShopClientInterface {
	return harryShop.New("https://api.jsonbin.io/b", 5)
}

func GetHarryBook(p productUsecase.ProductUsecaseInterface, h harryShop.HarryShopClientInterface) {
	datas, err := h.GetHarryBook()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	for _, data := range datas {

		err = p.Save(ctx, data)

		if err != nil && err != productDomain.ErrorProductExist {
			panic(err)
		}
	}
}
