package service

import (
	"context"
	"ubankApi/domain"
	"ubankApi/repository"
)

type TransactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepo}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, transaction *domain.Transaction) error {
	return s.transactionRepo.CreateTransaction(ctx, transaction)
}

// Autres méthodes de service pour la gestion des transactions (recherche, mise à jour, suppression, etc.) peuvent être ajoutées ici.
