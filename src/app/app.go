package app

import (
	"context"
	"github.com/labstack/echo"
	"github.com/oommi04/shibabookbackend/src/configs"
	"github.com/oommi04/shibabookbackend/src/drivers/echoDriver"
	"github.com/oommi04/shibabookbackend/src/external/google"
	"github.com/oommi04/shibabookbackend/src/external/lineBot"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

func SetupHttp(c *configs.Configs) *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	middlewareEcho := echoDriver.InitMiddleware()
	e.Use(middlewareEcho.CORS)

	return e
}


func SetupMongo() *mongo.Database{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	//defer client.Disconnect(ctx)
	database := client.Database("shibaBookShop")
	return database
}