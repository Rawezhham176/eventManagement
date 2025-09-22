package handlers

import (
	"eventManagement/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not fetch event. Try again later": err.Error()})
		return
	}

	event, err := model.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not fetch event. Try again later": err.Error()})
		return
	}

	err = event.RegisterEvent(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not register event. Try again later": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"Event": event})
}

func CancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not fetch event. Try again later": err.Error()})
		return
	}

	var event model.Event
	event.ID = eventId

	err = event.CancelEvent(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not register event. Try again later": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"Event": event})
}

func GetEventAttendees(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not fetch event. Try again later": err.Error()})
		return
	}

	event, err := model.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not fetch event. Try again later": err.Error()})
		return
	}

	registrationList, err := event.GetRegistrationList(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not register event. Try again later": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"Registration and Attendees List": registrationList})
}

func GetRegistrationsByUserId(c *gin.Context) {
	userId := c.GetInt64("userId")

	registrations, err := model.GetRegistrationsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not register event. Try again later": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"Registrations from user": registrations})
}
