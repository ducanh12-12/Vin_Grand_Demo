package controllers

import (
	"base-go/application/auth"
	"base-go/common/logger"
	validate "base-go/gateway/http/validator/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authInteractor auth.AuthInteractor
}

func (controller *AuthController) Mount(e *echo.Echo) {
	g := e.Group("/auth")
	g.POST("/login", controller.Login)
}

func NewAuthController(authInteractor auth.AuthInteractor) *AuthController {
	return &AuthController{authInteractor}
}

func (controller *AuthController) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	phone_number := c.FormValue("phone_number")
	var token string
	var err error
	err = validate.ValidateLogin(password, username, phone_number)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	if phone_number == "" {
		token, err = controller.authInteractor.Login(c.Request().Context(), username, "", password)
	} else {
		token, err = controller.authInteractor.Login(c.Request().Context(), "", phone_number, password)
	}
	logger.Info("Login")
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusUnauthorized, err.Error())
		return nil
	}
	c.JSON(http.StatusOK, token)
	return nil
}
