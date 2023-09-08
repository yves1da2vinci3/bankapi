package domain

import "time"

type Transaction struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

// Vous pouvez ajouter d'autres méthodes ou fonctions liées à la gestion des transactions ici.
