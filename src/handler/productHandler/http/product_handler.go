package http

import (
	"context"
	"github.com/labstack/echo"
	"github.com/oommi04/shibabookbackend/src/domains/productDomain"
	"github.com/oommi04/shibabookbackend/src/usecase/productUsecase"
	"github.com/oommi04/shibabookbackend/src/utils/errorStatus"
	"net/http"
)

type responseError struct {
	Message string `json:"message"`
}

type body struct {
	ProductId string                `json:"productId,omitempty"`
	Product   productDomain.Product `json:"product,omitempty"`
}

type productHandler struct {
	usecase productUsecase.ProductUsecaseInterface
}

func Init(e *echo.Echo, u productUsecase.ProductUsecaseInterface) {
	handler := productHandler{
		u,
	}
	e.GET("/product", handler.List)
	e.POST("/product", handler.Save)
}

func (h *productHandler) List(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	resp, err := h.usecase.List(ctx)

	if err != nil {
		return c.JSON(errorStatus.GetStatusCode(err), responseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *productHandler) Save(c echo.Context) error {
	var data productDomain.Product
	var body body
	err := c.Bind(&body)
	data = body.Product

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = h.usecase.Save(ctx, &data)

	if err != nil {
		return c.JSON(errorStatus.GetStatusCode(err), responseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, data)
}
