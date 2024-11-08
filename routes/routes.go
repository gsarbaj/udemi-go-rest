package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", GetHello)

	server.GET("/events", getEvents)
	server.POST("/", createEvent)

	server.GET("/actions", GetActions)
	server.GET("/actions/:id", GetAction)
	server.POST("/actions", CreateAction)
	server.PUT("/actions/:id", UpdateAction)
	server.DELETE("actions/:id", DeleteAction)
	server.POST("/signup", signUpHandler)
	server.POST("/login")
}
