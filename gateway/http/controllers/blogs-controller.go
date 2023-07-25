package controllers

import (
	"base-go/application/blogs"
	"base-go/common/logger"
	"base-go/gateway/http/presenter"
	"base-go/gateway/http/validator/blog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BlogsController struct {
	blogsInteractor blogs.BlogsInteractor
}

func NewBlogsController(blogsInteractor blogs.BlogsInteractor) *BlogsController {
	return &BlogsController{blogsInteractor}
}

func (controller *BlogsController) Mount(e *echo.Echo) {
	g := e.Group("/blogs")
	g.GET("/:id", controller.GetBlog)
	g.GET("", controller.GetBlogs)
	g.POST("", controller.AddBlog)
}

func (controller *BlogsController) GetBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// should validate ipt here
	logger.Info("GetBlog input: id=%d", id)
	blog, err := controller.blogsInteractor.GetBlog(c.Request().Context(), id)
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err)
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, blog)
	return nil
}
func (controller *BlogsController) GetBlogs(c echo.Context) error {
	logger.Info("Get list blog ")
	blog, err := controller.blogsInteractor.GetBlogs(c.Request().Context())
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err)
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, blog)
	return nil
}
func (controller *BlogsController) AddBlog(c echo.Context) error {
	blogIpt := blogs.AddBlogIpt{}
	logger.Info("AddBlog input: %+v", blogIpt)
	err := c.Bind(&blogIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil
	}
	err = blog.ValidateBlog(blogIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil
	}
	logger.Info("AddBlog input: %+v", c)
	newblog, err := controller.blogsInteractor.CreateBlog(c.Request().Context(), blogIpt)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
		return nil
	}
	blog := presenter.Blog(newblog)
	c.JSON(http.StatusOK, blog)
	return nil
}
