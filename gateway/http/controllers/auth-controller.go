package controllers

import (
	"base-go/application/auth"
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
	token, err := controller.authInteractor.Login(c.Request().Context(), username, password)
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err)
		return nil
	}
	c.JSON(http.StatusOK, token)
	return nil
}
