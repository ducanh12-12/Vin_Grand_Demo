package controllers

import (
	"base-go/application/events"
	"base-go/common/logger"
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
	g := e.Group("/events")
	g.GET("/:id", controller.GetEvent)
	g.GET("", controller.GetEvents)
	g.POST("", controller.AddEvent)
}

func (controller *EventsController) GetEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// should validate ipt here
	logger.Info("GetEvent input: id=%d", id)
	event, err := controller.eventsInteractor.GetEvent(c.Request().Context(), id)
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err)
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
		c.JSON(http.StatusBadGateway, err)
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
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil
	}
	err = event.ValidateEvent(eventIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil
	}
	logger.Info("AddEvent input: %+v", c)
	newEvent, err := controller.eventsInteractor.CreateEvent(c.Request().Context(), eventIpt)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
		return nil
	}
	event := presenter.Event(newEvent)
	c.JSON(http.StatusOK, event)
	return nil
}
