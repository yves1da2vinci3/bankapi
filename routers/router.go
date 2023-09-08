package routers

import (
	"ubankApi/security"

	"github.com/gofiber/fiber/v2"
)

func SetupApp(db *gorm.DB, jwtUtil security.JWTUtil) *fiber.App {
	app := fiber.New()

	// Middleware global (facultatif)
	// app.Use(middleware.GlobalMiddleware())

	// Configurer les routes pour chaque entité
	SetupUserRoutes(app, db, jwtUtil)
	SetupAuthRoutes(app, db, jwtUtil)
	SetupTransactionRoutes(app, db, jwtUtil)

	// ...

	return app
}
