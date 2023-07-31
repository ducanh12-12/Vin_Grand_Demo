package controllers

import (
	"base-go/application/events"
	"base-go/common/logger"
	"base-go/gateway/http/middlewares"
	"base-go/gateway/http/presenter"
	"base-go/gateway/http/validator/event"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EventsController struct {
	eventsInteractor events.EventsInteractor
}

func NewEventsController(eventsInteractor events.EventsInteractor) *EventsController {
	return &EventsController{eventsInteractor}
}

func (controller *EventsController) Mount(e *echo.Echo) {
	g := e.Group("/api/events", middlewares.MiddlewareJWT)
	g.GET("/:id", controller.GetEvent)
	g.GET("", controller.GetEvents)
	g.POST("", controller.AddEvent)
	g.PUT("/:id", controller.Update)
	g.DELETE("/:id", controller.Delete)
}

func (controller *EventsController) GetEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// should validate ipt here
	logger.Info("GetEvent input: id=%d", id)
	event, err := controller.eventsInteractor.GetEvent(c.Request().Context(), id)
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, event)
	return nil
}
func (controller *EventsController) GetEvents(c echo.Context) error {
	logger.Info("Get list event ")
	event, err := controller.eventsInteractor.GetEvents(c.Request().Context())
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, event)
	return nil
}
func (controller *EventsController) AddEvent(c echo.Context) error {
	eventIpt := events.AddEventIpt{}
	logger.Info("AddEvent input: %+v", eventIpt)
	err := c.Bind(&eventIpt)
	// if eventIpt.StartTime != "" {
	// 	layout := "02/01/2006"
	// 	t, err := time.Parse(layout, eventIpt.StartTime)
	// 	if err != nil {
	// 		fmt.Println("Failed to parse StartTime:", err.Error())
	// 	} else {
	// 		eventIpt.StartTime := &t
	// 	}
	// }

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	err = event.ValidateEvent(eventIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	logger.Info("AddEvent input: %+v", c)
	newEvent, err := controller.eventsInteractor.CreateEvent(c.Request().Context(), eventIpt)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	event := presenter.Event(newEvent)
	c.JSON(http.StatusOK, event)
	return nil
}
func (controller *EventsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	eventIpt := events.UpdateEventIpt{}
	logger.Info("AddBlog input: %+v", eventIpt)
	err := c.Bind(&eventIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	err = event.ValidateUpdateEvent(eventIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	logger.Info("AddBlog input: %+v", c)
	newblog, err := controller.eventsInteractor.Update(c.Request().Context(), eventIpt, id)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	blog := presenter.Event(newblog)
	c.JSON(http.StatusOK, blog)
	return nil
}
func (controller *EventsController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// should validate ipt here
	logger.Info("Delete input: id=%d", id)
	user, err := controller.eventsInteractor.Delete(c.Request().Context(), id)
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err.Error())
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, user)
	return nil
}
