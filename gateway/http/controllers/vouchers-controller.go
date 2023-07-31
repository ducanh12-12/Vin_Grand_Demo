package controllers

import (
	"base-go/application/vouchers"
	"base-go/common/logger"
	"base-go/gateway/http/middlewares"
	"base-go/gateway/http/presenter"
	"base-go/gateway/http/validator/voucher"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type VouchersController struct {
	vouchersInteractor vouchers.VouchersInteractor
}

func NewVouchersController(vouchersInteractor vouchers.VouchersInteractor) *VouchersController {
	return &VouchersController{vouchersInteractor}
}

func (controller *VouchersController) Mount(e *echo.Echo) {
	g := e.Group("/api/vouchers", middlewares.MiddlewareJWT)
	g.GET("/:id", controller.GetVoucher)
	g.GET("", controller.GetVouchers)
	g.POST("", controller.AddVoucher)
	g.PUT("/:id", controller.Update)
	g.DELETE("/:id", controller.Delete)
}

func (controller *VouchersController) GetVoucher(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// should validate ipt here
	logger.Info("GetVoucher input: id=%d", id)
	voucher, err := controller.vouchersInteractor.GetVoucher(c.Request().Context(), id)
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, voucher)
	return nil
}
func (controller *VouchersController) GetVouchers(c echo.Context) error {
	logger.Info("Get list voucher ")
	voucher, err := controller.vouchersInteractor.GetVouchers(c.Request().Context())
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, voucher)
	return nil
}
func (controller *VouchersController) AddVoucher(c echo.Context) error {
	voucherIpt := vouchers.AddVoucherIpt{}
	logger.Info("AddVoucher input: %+v", voucherIpt)
	err := c.Bind(&voucherIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	err = voucher.ValidateVoucher(voucherIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	logger.Info("AddVoucher input: %+v", c)
	newVoucher, err := controller.vouchersInteractor.CreateVoucher(c.Request().Context(), voucherIpt)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	voucher := presenter.Voucher(newVoucher)
	c.JSON(http.StatusOK, voucher)
	return nil
}
func (controller *VouchersController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	eventIpt := vouchers.UpdateVoucherIpt{}
	logger.Info("AddBlog input: %+v", eventIpt)
	err := c.Bind(&eventIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	err = voucher.ValidateUpdateVoucher(eventIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	logger.Info("AddBlog input: %+v", c)
	newblog, err := controller.vouchersInteractor.Update(c.Request().Context(), eventIpt, id)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	blog := presenter.Voucher(newblog)
	c.JSON(http.StatusOK, blog)
	return nil
}
func (controller *VouchersController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// should validate ipt here
	logger.Info("Delete input: id=%d", id)
	user, err := controller.vouchersInteractor.Delete(c.Request().Context(), id)
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, user)
	return nil
}
