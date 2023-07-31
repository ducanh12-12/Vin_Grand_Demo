package controllers

import (
	"base-go/application/invoices"
	"base-go/common/logger"
	"base-go/gateway/http/middlewares"
	"base-go/gateway/http/presenter"
	"base-go/gateway/http/validator/invoice"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type InvoicesController struct {
	invoicesInteractor invoices.InvoicesInteractor
}

func NewInvoicesController(invoicesInteractor invoices.InvoicesInteractor) *InvoicesController {
	return &InvoicesController{invoicesInteractor}
}

func (controller *InvoicesController) Mount(e *echo.Echo) {
	g := e.Group("/api/invoices", middlewares.MiddlewareJWT)
	g.GET("/:id", controller.Get)
	g.GET("", controller.Gets)
	g.POST("", controller.Create)
	g.PUT("/:id", controller.Update)
	g.DELETE("/:id", controller.Delete)
}

func (controller *InvoicesController) Get(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// should validate ipt here
	logger.Info("Get input: id=%d", id)
	invoice, err := controller.invoicesInteractor.Retrieve(c.Request().Context(), id)
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, invoice)
	return nil
}
func (controller *InvoicesController) Gets(c echo.Context) error {
	logger.Info("Get list invoice ")
	invoice, err := controller.invoicesInteractor.List(c.Request().Context())
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, invoice)
	return nil
}
func ConvertPriceToPoint(point float64) int {
	const one_point = 40
	return int(math.Ceil(point / one_point))
}
func (controller *InvoicesController) Create(c echo.Context) error {
	invoiceIpt := invoices.AddInvoiceIpt{}
	logger.Info("Create input: %+v", invoiceIpt)
	err := c.Bind(&invoiceIpt)
	invoiceIpt.Point = ConvertPriceToPoint(invoiceIpt.Price)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	err = invoice.ValidateInvoice(invoiceIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	logger.Info("Create invoice: %+v", c)
	newinvoice, err := controller.invoicesInteractor.Create(c.Request().Context(), invoiceIpt)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	invoice := presenter.Invoice(newinvoice)
	c.JSON(http.StatusOK, invoice)
	return nil
}
func (controller *InvoicesController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	eventIpt := invoices.UpdateInvoiceIpt{}
	logger.Info("AddBlog input: %+v", eventIpt)
	err := c.Bind(&eventIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	err = invoice.ValidateUpdateInvoice(eventIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	logger.Info("AddBlog input: %+v", c)
	newblog, err := controller.invoicesInteractor.Update(c.Request().Context(), id, eventIpt)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	blog := presenter.Invoice(newblog)
	c.JSON(http.StatusOK, blog)
	return nil
}
func (controller *InvoicesController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// should validate ipt here
	logger.Info("Delete input: id=%d", id)
	user, err := controller.invoicesInteractor.Delete(c.Request().Context(), id)
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, user)
	return nil
}
