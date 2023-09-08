package logic

import (
	"ubankApi/domain"
)

type ValidationLogic struct{}

func NewValidationLogic() *ValidationLogic {
	return &ValidationLogic{}
}

func (l *ValidationLogic) ValidateUser(user *domain.User) error {
	// Effectuez des validations sur l'entité User ici en utilisant des règles de validation spécifiques
	// Retournez une erreur en cas de validation échouée, sinon retournez nil.
}
