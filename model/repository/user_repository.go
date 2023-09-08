package repository

import (
	"context"
	"log"
	"ubankApi/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		log.Println("Failed to create user:", result.Error)
		return result.Error
	}
	return nil
}

// Autres méthodes pour la gestion des utilisateurs (recherche, mise à jour, suppression, etc.) peuvent être ajoutées ici.
