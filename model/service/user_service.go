package service

import (
	"context"
	"ubankApi/domain"
	"ubankApi/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	return s.userRepo.CreateUser(ctx, user)
}

// Autres méthodes de service pour la gestion des utilisateurs (recherche, mise à jour, suppression, etc.) peuvent être ajoutées ici.
