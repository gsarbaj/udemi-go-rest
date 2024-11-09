package routes

import (
	"github.com/gin-gonic/gin"
	"imta.icu/rest/models"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	actionId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse action id"})
		return
	}

	action, err := models.GetActionByID(actionId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find action"})
		return
	}

	err = action.RegisterAction(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not register action"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered action"})

}

func cancelRegistration(context *gin.Context) {}
