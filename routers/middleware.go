package routers

import (
	"ubankApi/security"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(jwtUtil security.JWTUtil) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Vérifier le jeton JWT et définir l'utilisateur dans le contexte si l'authentification réussit
		user, err := jwtUtil.VerifyToken(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		c.Locals("user", user)
		return c.Next()
	}
}
