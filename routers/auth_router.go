package routers

import (
	"ubankApi/controllers"
	"ubankApi/security"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, db *gorm.DB, jwtUtil security.JWTUtil) {
	// Routes d'authentification
	authoute := app.Group("/auth")
	authoute.Post("/login", controllers.Login(db, jwtUtil))
	authoute.Post("/register", controllers.Register(db))

	// ...
}
