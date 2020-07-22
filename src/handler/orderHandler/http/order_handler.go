package http

import (
	"context"
	"github.com/labstack/echo"
	"github.com/oommi04/shibabookbackend/src/domains/orderDomain"
	"github.com/oommi04/shibabookbackend/src/usecase/orderUsecase"
	"github.com/oommi04/shibabookbackend/src/utils/errorStatus"
	"net/http"
)

type responseError struct {
	Message string `json:"message"`
}

type body struct {
	OrderID string            `json:"orderId,omitempty"`
	Order   orderDomain.Order `json:"order,omitempty"`
}

type orderHandler struct {
	usecase orderUsecase.OrderUsecaseInterface
}

func Init(e *echo.Echo, u orderUsecase.OrderUsecaseInterface) {
	handler := orderHandler{
		u,
	}
	e.GET("/order", handler.List)
	e.POST("/order", handler.Save)
	e.POST("/order/checkOut", handler.Checkout)
}

func (h *orderHandler) List(c echo.Context) error {
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

func (h *orderHandler) Save(c echo.Context) error {
	var data orderDomain.Order
	var body body
	err := c.Bind(&body)
	data = body.Order

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

func (h *orderHandler) Checkout(c echo.Context) error {
	var body body
	err := c.Bind(&body)
	orderId := body.OrderID

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	resp, err := h.usecase.CheckOut(ctx, orderId)

	if err != nil {
		return c.JSON(errorStatus.GetStatusCode(err), responseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, resp)
}
