package routes

import (
	"github.com/gin-gonic/gin"
	"imta.icu/rest/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", GetHello)

	server.GET("/events", getEvents)
	server.POST("/", createEvent)

	server.GET("/actions", GetActions)
	server.GET("/actions/:id", GetAction)

	//protected routes

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/actions", CreateAction)
	authenticated.PUT("/actions/:id", UpdateAction)
	authenticated.DELETE("actions/:id", DeleteAction)
	authenticated.POST("actions/:id/register", registerForEvent)
	authenticated.DELETE("actions/:id/register")

	server.POST("/signup", signUpHandler)
	server.POST("/login", login)
}
