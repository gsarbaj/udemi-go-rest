package routes

import (
	"github.com/gin-gonic/gin"
	"imta.icu/rest/models"
	"net/http"
)

func signUpHandler(context *gin.Context) {
	var user models.User

	err := context.ShouldBind(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user.ID})
}
