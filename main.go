package main

import (
	"github.com/tkhamsila/shibabookbackend/src/app"
	"github.com/tkhamsila/shibabookbackend/src/configs"
	customerHandler "github.com/tkhamsila/shibabookbackend/src/handler/customerHandler/http"
	orderHandler "github.com/tkhamsila/shibabookbackend/src/handler/orderHandler/http"
	productHandler "github.com/tkhamsila/shibabookbackend/src/handler/productHandler/http"
	customerRepo "github.com/tkhamsila/shibabookbackend/src/repository/customerRepository/mongo"
	invoiceRepo "github.com/tkhamsila/shibabookbackend/src/repository/invoiceRepository/mongo"
	orderRepo "github.com/tkhamsila/shibabookbackend/src/repository/orderRepository/mongo"
	productRepo "github.com/tkhamsila/shibabookbackend/src/repository/productRepository/mongo"
	"github.com/tkhamsila/shibabookbackend/src/usecase/customerUsecase"
	"github.com/tkhamsila/shibabookbackend/src/usecase/orderUsecase"
	"github.com/tkhamsila/shibabookbackend/src/usecase/productUsecase"
	"time"
)

func main() {
	cfg := configs.New()

	e := app.SetupHttp(cfg)

	mdb := app.SetupMongo()
	timeOutContext := 5 * time.Second

	productRepoInstance := productRepo.NewProductRepository(mdb)
	productUsecaseInstance := productUsecase.NewProductUsecase(productRepoInstance, timeOutContext)
	productHandler.Init(e,productUsecaseInstance)

	customerRepoInstance := customerRepo.NewCustomerRepository(mdb)
	customerUsecaseInstance := customerUsecase.NewProductUsecase(customerRepoInstance,timeOutContext)
	customerHandler.Init(e,customerUsecaseInstance)

	invoiceRepoInstance := invoiceRepo.NewInvoiceRepository(mdb)

	orderRepoInstance := orderRepo.NewProductRepository(mdb)
	orderUsecaseInstance := orderUsecase.NewProductUsecase(orderRepoInstance,productRepoInstance, invoiceRepoInstance,customerUsecaseInstance ,timeOutContext)
	orderHandler.Init(e,orderUsecaseInstance)

	hs := app.SetUpHarryShop()
	app.GetHarryBook(productUsecaseInstance, hs)

	e.Logger.Fatal(e.Start(":" + cfg.PORT))

}