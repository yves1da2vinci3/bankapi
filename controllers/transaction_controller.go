package controllers

import (
	"context"
	"ubankApi/domain"
	"ubankApi/service"

	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *TransactionController {
	return &TransactionController{transactionService}
}

func (c *TransactionController) CreateTransaction(ctx *fiber.Ctx) error {
	transaction := new(domain.Transaction)

	if err := ctx.BodyParser(transaction); err != nil {
		return err
	}

	if err := c.transactionService.CreateTransaction(context.TODO(), transaction); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(transaction)
}

// Autres méthodes de contrôleur pour la gestion des transactions (recherche, mise à jour, suppression, etc.) peuvent être ajoutées ici.
