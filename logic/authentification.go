package logic

import (
	"context"
	"ubankApi/repository"
	"ubankApi/security"
)

type AuthentificationLogic struct {
	userRepository repository.UserRepository
	jwtUtil        security.JWTUtil
}

func NewAuthentificationLogic(userRepository repository.UserRepository, jwtUtil security.JWTUtil) *AuthentificationLogic {
	return &AuthentificationLogic{userRepository, jwtUtil}
}

func (l *AuthentificationLogic) Authenticate(ctx context.Context, email, password string) (string, error) {
	// Vérifiez les informations d'identification de l'utilisateur ici en utilisant le userRepository
	// Si l'utilisateur est authentifié, générez un token JWT à renvoyer
	// Utilisez jwtUtil pour générer le token JWT
	// Retournez le token JWT en cas de succès, sinon renvoyez une erreur d'authentification.
}

// Autres méthodes d'authentification, telles que la réinitialisation du mot de passe, peuvent être ajoutées ici.
