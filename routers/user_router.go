package routers

import (
	"ubankApi/controllers"
	"ubankApi/security"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, db *gorm.DB, jwtUtil security.JWTUtil) {
	// Routes CRUD pour les utilisateurs
	userRoute := app.Group("/users")
	userRoute.Get("/", controllers.GetUsers(db))
	userRoute.Get("/:id", controllers.GetUser(db))
	userRoute.Put("/:id", controllers.UpdateUser(db))
	userRoute.Delete("/:id", controllers.DeleteUser(db))

	// ...
}
