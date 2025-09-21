package routes

import (
	"eventManagement/handlers"
	"eventManagement/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// public endpoints
	server.GET("/events", handlers.GetEvents)
	server.GET("/event/:id", handlers.GetEventById)
	server.GET("/events/search", handlers.SearchEventsByNameOrLocation)
	server.GET("/events/categories", handlers.GetEventByCategory)
	server.GET("/events/upcoming", handlers.GetUpcomingEvents)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/event", handlers.CreateEvent)
	authenticated.PUT("/event/:id", handlers.UpdateEvent)
	authenticated.DELETE("/event/:id", handlers.DeleteEvent)

	authenticated.POST("/events/:id/register", handlers.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", handlers.CancelRegistration)

	server.POST("/user/signup", handlers.CreateUser)
	server.POST("/user/login", handlers.LoginUser)
}
