package routers

import (
	"ubankApi/controllers"
	"ubankApi/security"

	"github.com/gofiber/fiber/v2"
)

func SetupTransactionRoutes(app *fiber.App, db *gorm.DB, jwtUtil security.JWTUtil) {
	// Routes CRUD pour les transactions
	transactionRoute := app.Group("/transactions")
	transactionRoute.Get("/", controllers.GetTransactions(db))
	transactionRoute.Get("/:id", controllers.GetTransaction(db))
	transactionRoute.Post("/", controllers.CreateTransaction(db))
	transactionRoute.Put("/:id", controllers.UpdateTransaction(db))
	transactionRoute.Delete("/:id", controllers.DeleteTransaction(db))

	// ...
}
