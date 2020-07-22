package http

import (
	"context"
	"github.com/labstack/echo"
	"github.com/oommi04/shibabookbackend/src/domains/customerDomain"
	"github.com/oommi04/shibabookbackend/src/usecase/customerUsecase"
	"github.com/oommi04/shibabookbackend/src/utils/errorStatus"
	"net/http"
)

type responseError struct {
	Message string `json:"message"`
}

type body struct {
	Customer customerDomain.Customer `json:"customer,omitempty"`
}

type customerHandler struct {
	usecase customerUsecase.CustomerUsecaseInterface
}

func Init(e *echo.Echo, u customerUsecase.CustomerUsecaseInterface) {
	handler := customerHandler{
		u,
	}
	e.POST("/customer", handler.Save)
}

func (h *customerHandler) Save(c echo.Context) error {
	var data customerDomain.Customer
	var body body
	err := c.Bind(&body)
	data = body.Customer

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
