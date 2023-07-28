package controllers

import (
	"base-go/application/users"
	"base-go/common/logger"
	"base-go/gateway/http/middlewares"
	"base-go/gateway/http/presenter"
	"base-go/gateway/http/validator/user"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UsersController struct {
	usersInteractor users.UsersInteractor
}

func NewUsersController(usersInteractor users.UsersInteractor) *UsersController {
	return &UsersController{usersInteractor}
}

func (controller *UsersController) Mount(e *echo.Echo) {
	g := e.Group("/api/users", middlewares.MiddlewareJWT)
	g.GET("/:id", controller.GetUser)
	g.PUT("/add_point/:id", controller.AddPoint)
	g.GET("", controller.GetUsers)
	g.POST("", controller.Adduser)
}

// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "UserID" Format(int)
// @Success 201 object model.User
// @Router /users/{id} [get]
func (controller *UsersController) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// should validate ipt here
	logger.Info("Getuser input: id=%d", id)
	user, err := controller.usersInteractor.GetUser(c.Request().Context(), id)
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err)
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, user)
	return nil
}
func (controller *UsersController) GetUsers(c echo.Context) error {
	logger.Info("Get list user ")
	user, err := controller.usersInteractor.GetUsers(c.Request().Context())
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err)
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, user)
	return nil
}
func convertToPointer(str string) *string {
	return &str
}
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func (controller *UsersController) Adduser(c echo.Context) error {
	pasworDefault := os.Getenv("DEFAULT_PASSWORD")
	userIpt := users.AddUserIpt{}
	PasswordHash, err := Hash(pasworDefault)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil
	}
	logger.Info("Adduser input: %+v", userIpt)
	err = c.Bind(&userIpt)
	userExit, err := controller.usersInteractor.GetUserByUserName(c.Request().Context(), userIpt.UserName)
	if userExit.UserName != "" {
		c.JSON(http.StatusBadRequest, "User exit")
		return nil
	}
	err = nil
	userIpt.Password = PasswordHash
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	err = user.ValidateUser(userIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	logger.Info("Adduser input: %+v", c)
	newuser, err := controller.usersInteractor.AddUser(c.Request().Context(), userIpt)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
		return nil
	}
	user := presenter.User(newuser)
	c.JSON(http.StatusOK, user)
	return nil
}
func (controller *UsersController) AddPoint(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	userIpt := users.AddPoint{}
	err := c.Bind(&userIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	err = user.ValidatePoint(userIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	logger.Info("Add point user input: %+v", c)
	newuser, err := controller.usersInteractor.AddPointUser(c.Request().Context(), userIpt, id)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	user := presenter.User(newuser)
	c.JSON(http.StatusOK, user)
	return nil
}
