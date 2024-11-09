package routes

import (
	"github.com/gin-gonic/gin"
	"imta.icu/rest/models"
	"net/http"
	"strconv"
)

func GetActions(context *gin.Context) {
	actions, err := models.GetAllActions()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"actions": actions})
}

func GetAction(context *gin.Context) {
	actionId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	action, err := models.GetActionByID(actionId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, action)
}

func CreateAction(context *gin.Context) {

	var action models.Action
	err := context.ShouldBindJSON(&action)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userId := context.GetInt64("userId")
	action.UserID = userId

	err = action.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Action created", "action": action})

}

func UpdateAction(context *gin.Context) {
	actionId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	action, err := models.GetActionByID(actionId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if action.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "You can't modify your own action"})
		return
	}

	var updatedAction models.Action
	err = context.ShouldBindJSON(&updatedAction)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAction.ID = actionId

	err = updatedAction.UpdateAction()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Action updated", "action": updatedAction})
}

func DeleteAction(context *gin.Context) {
	actionId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	action, err := models.GetActionByID(actionId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if action.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "You can't delete not your own action"})
		return
	}

	err = action.DeleteAction()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Action deleted", "action": action})

}
