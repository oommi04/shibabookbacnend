package main

import (
	"github.com/oommi04/shibabookbackend/src/app"
	"github.com/oommi04/shibabookbackend/src/configs"
	customerHandler "github.com/oommi04/shibabookbackend/src/handler/customerHandler/http"
	orderHandler "github.com/oommi04/shibabookbackend/src/handler/orderHandler/http"
	productHandler "github.com/oommi04/shibabookbackend/src/handler/productHandler/http"
	customerRepo "github.com/oommi04/shibabookbackend/src/repository/customerRepository/mongo"
	invoiceRepo "github.com/oommi04/shibabookbackend/src/repository/invoiceRepository/mongo"
	orderRepo "github.com/oommi04/shibabookbackend/src/repository/orderRepository/mongo"
	productRepo "github.com/oommi04/shibabookbackend/src/repository/productRepository/mongo"
	"github.com/oommi04/shibabookbackend/src/usecase/customerUsecase"
	"github.com/oommi04/shibabookbackend/src/usecase/orderUsecase"
	"github.com/oommi04/shibabookbackend/src/usecase/productUsecase"
	"time"
)

func main() {
	cfg := configs.New()

	e := app.SetupHttp(cfg)

	mdb := app.SetupMongo()
	timeOutContext := 5 * time.Second

	productRepoInstance := productRepo.New(mdb)
	productUsecaseInstance := productUsecase.New(productRepoInstance, timeOutContext)
	productHandler.Init(e,productUsecaseInstance)

	customerRepoInstance := customerRepo.New(mdb)
	customerUsecaseInstance := customerUsecase.New(customerRepoInstance,timeOutContext)
	customerHandler.Init(e,customerUsecaseInstance)

	invoiceRepoInstance := invoiceRepo.New(mdb)

	orderRepoInstance := orderRepo.New(mdb)
	orderUsecaseInstance := orderUsecase.New(orderRepoInstance,productRepoInstance, invoiceRepoInstance,customerUsecaseInstance ,timeOutContext)
	orderHandler.Init(e,orderUsecaseInstance)

	hs := app.SetUpHarryShop()
	app.GetHarryBook(productUsecaseInstance, hs)

	e.Logger.Fatal(e.Start(":" + cfg.PORT))

}