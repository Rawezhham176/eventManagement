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

	// Event Management
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/event", handlers.CreateEvent)
	authenticated.PUT("/event/:id", handlers.UpdateEvent)
	authenticated.DELETE("/event/:id", handlers.DeleteEvent)

	// Event Participation
	authenticated.POST("/events/:id/register", handlers.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", handlers.CancelRegistration)
	authenticated.GET("/events/:id/attendees", handlers.GetEventAttendees)
	authenticated.GET("/my-events", handlers.GetEventsByUserId)
	authenticated.GET("/my-registrations", handlers.GetRegistrationsByUserId) // NEW: Meine Anmeldungen

	// Public User Endpoints
	server.POST("/user/signup", handlers.CreateUser)
	server.POST("/user/login", handlers.LoginUser)
	server.POST("/user/forgot-password", handlers.ForgotPassword)
	server.POST("/user/reset-password", handlers.ResetPassword)

	// User Profile Management
	/*	authenticated.GET("/user/profile", getUserProfile)          // NEW: Profil abrufen
		authenticated.PUT("/user/profile", updateUserProfile)       // NEW: Profil bearbeiten
		authenticated.POST("/user/change-password", changePassword) // NEW: Passwort ändern
		authenticated.DELETE("/user/account", deleteUserAccount)    // NEW: Account löschen
		authenticated.POST("/user/upload-avatar", uploadAvatar)     // NEW: Profilbild hochladen

		// Comments & Reviews
		authenticated.POST("/events/:id/comments", createComment)   // NEW: Kommentar erstellen
		authenticated.GET("/events/:id/comments", getEventComments) // NEW: Kommentare abrufen
		authenticated.PUT("/comments/:id", updateComment)           // NEW: Kommentar bearbeiten
		authenticated.DELETE("/comments/:id", deleteComment)        // NEW: Kommentar löschen
		authenticated.POST("/events/:id/review", createReview)      // NEW: Bewertung erstellen
		authenticated.GET("/events/:id/reviews", getEventReviews)   // NEW: Bewertungen abrufen

		// Favorites & Wishlist
		authenticated.POST("/events/:id/favorite", addToFavorites)        // NEW: Zu Favoriten hinzufügen
		authenticated.DELETE("/events/:id/favorite", removeFromFavorites) // NEW: Aus Favoriten entfernen
		authenticated.GET("/user/favorites", getUserFavorites)            // NEW: Favoriten abrufen

		// Notifications
		authenticated.GET("/notifications", getNotifications)                // NEW: Benachrichtigungen abrufen
		authenticated.PUT("/notifications/:id/read", markNotificationAsRead) // NEW: Als gelesen markieren
		authenticated.DELETE("/notifications/:id", deleteNotification)       // NEW: Benachrichtigung löschen

		// Statistics (for event organizers)
		authenticated.GET("/events/:id/statistics", getEventStatistics)    // NEW: Event-Statistiken
		authenticated.GET("/user/organizer-stats", getOrganizerStatistics) // NEW: Organisator-Statistiken

		// Admin Routes (mit Admin-Middleware)
		admin := server.Group("/admin")
		admin.Use(middleware.Authenticate, middleware.RequireAdmin) // Du brauchst noch Admin-Middleware
		admin.GET("/users", getAllUsers)                            // NEW: Alle User (Admin)
		admin.PUT("/users/:id/status", updateUserStatus)            // NEW: User sperren/entsperren
		admin.GET("/events/pending", getPendingEvents)              // NEW: Events zur Überprüfung
		admin.PUT("/events/:id/approve", approveEvent)              // NEW: Event genehmigen
		admin.GET("/reports", getReports)                           // NEW: Gemeldete Inhalte
		admin.PUT("/reports/:id/resolve", resolveReport)*/
}
