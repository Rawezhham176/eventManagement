package handlers

import (
	"eventManagement/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	events, err := model.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func CreateEvent(c *gin.Context) {
	var newEvent model.Event
	err := c.ShouldBindJSON(&newEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEvent.UserID = c.GetInt64("userId")

	err = newEvent.Save()
	if err != nil {
		panic(fmt.Sprintf("create events table error: %s", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"event created": newEvent})
}

func GetEventById(c *gin.Context) {
	parseInt, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not fetch events. Try again later": err.Error()})
		return
	}

	userId := c.GetInt64("userId")
	event, err := model.GetEventById(parseInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not fetch event": err.Error()})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You do not have permission to access this event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func UpdateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not fetch event. Try again later": err.Error()})
		return
	}

	userId := c.GetInt64("userId")
	event, err := model.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not fetch event": err.Error()})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You do not have permission to access this event"})
		return
	}

	var newEvent model.Event
	err = c.ShouldBindJSON(&newEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEvent.ID = eventId
	err = newEvent.UpdateEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not fetch event. Try again later": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event updated": newEvent})
}

func DeleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not fetch event. Try again later": err.Error()})
		return
	}

	userId := c.GetInt64("userId")
	event, err := model.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not fetch event": err.Error()})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You do not have permission to access this event"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not fetch event": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"event deleted": event})
}

func SearchEventsByNameOrLocation(c *gin.Context) {
	searchedName := c.Param("name")
	searchedLocation := c.Param("location")

	event, err := model.GetEventByNamOrLocation(searchedName, searchedLocation)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not find event": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func GetEventByCategory(c *gin.Context) {
	var body map[string]interface{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	category := body["category"].(string)

	event, err := model.GetEventsByCategory(category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not find event": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func GetUpcomingEvents(c *gin.Context) {
	var body map[string]interface{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fromQuery := body["from"].(string)
	toQuery := body["to"].(string)

	events, err := model.GetUpcomingEvents(fromQuery, toQuery)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not find event": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}
