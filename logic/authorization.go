package logic

import (
	"context"
	"ubankApi/repository"
)

type AuthorizationLogic struct {
	userRepository repository.UserRepository
}

func NewAuthorizationLogic(userRepository repository.UserRepository) *AuthorizationLogic {
	return &AuthorizationLogic{userRepository}
}

func (l *AuthorizationLogic) IsAuthorized(ctx context.Context, userID int, requiredRole string) (bool, error) {
	// Vérifiez si l'utilisateur avec l'ID userID a le rôle requiredRole
	// Utilisez le userRepository pour obtenir les informations de l'utilisateur et vérifier son rôle
	// Retournez true si l'utilisateur est autorisé, sinon retournez false
}
