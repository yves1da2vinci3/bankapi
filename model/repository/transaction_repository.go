package repository

import (
	"context"
	"log"
	"ubankApi/domain"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction *domain.Transaction) error {
	result := r.db.Create(transaction)
	if result.Error != nil {
		log.Println("Failed to create transaction:", result.Error)
		return result.Error
	}
	return nil
}

// Autres méthodes pour la gestion des transactions (recherche, mise à jour, suppression, etc.) peuvent être ajoutées ici.
