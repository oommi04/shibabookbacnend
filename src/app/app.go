package app

import (
	"context"
	"github.com/labstack/echo"
	"github.com/oommi04/shibabookbackend/src/configs"
	"github.com/oommi04/shibabookbackend/src/drivers/echoDriver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
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

func SetupMongo(ctx context.Context) (*mongo.Database, *mongo.Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	database := client.Database("shibaBookShop")
	return database, client
}
