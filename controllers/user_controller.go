package controllers

import (
	"context"
	"ubankApi/domain"
	"ubankApi/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	user := new(domain.User)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	if err := c.userService.CreateUser(context.TODO(), user); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

// Autres méthodes de contrôleur pour la gestion des utilisateurs (recherche, mise à jour, suppression, etc.) peuvent être ajoutées ici.
