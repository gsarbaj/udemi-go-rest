package routes

import (
	"github.com/gin-gonic/gin"
	"imta.icu/rest/models"
	"net/http"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch events", "error": err.Error()})
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

}
