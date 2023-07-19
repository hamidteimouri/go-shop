package httpserver

import (
	"github.com/go-playground/validator/v10"
	"github.com/hamidteimouri/go-shop/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ProductHandler struct {
	service   *domain.ProductService
	validator *validator.Validate
}

func RegisterProductHandler(e *echo.Echo, ctrl *domain.ProductService) {
	handler := &ProductHandler{
		service:   ctrl,
		validator: validator.New(),
	}

	e.GET("/products", handler.getProducts)
}

func (f *ProductHandler) getProducts(c echo.Context) error {
	req := ProductRequest{}
	err := c.Bind(&req)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error()
		return c.JSON(http.StatusUnprocessableEntity, &Response{Message: "something went wrong"})
	}

	productReq := &domain.ProductSearchRequest{
		Category: req.Category,
		PerPage:  req.PerPage,
		Page:     req.Page,
	}

	result, err := f.service.Products(c.Request().Context(), productReq)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("failed to get products")
		return c.JSON(http.StatusInternalServerError, &Response{Message: "something went wrong"})
	}

	return c.JSON(http.StatusOK, result)
}

type ProductRequest struct {
	Category string `json:"category" query:"category"`
	PerPage  int    `json:"per_page" query:"per_page"`
	Page     int    `json:"page" query:"page"`
}

type Response struct {
	Message string `json:"message"`
}
