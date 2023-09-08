package domain

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// Vous pouvez ajouter d'autres méthodes ou fonctions liées à la gestion des utilisateurs ici.
