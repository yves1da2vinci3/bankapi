package main

import (
	"ubankApi/config"
	"ubankApi/database"
	"ubankApi/routers"
	"ubankApi/security"
)

func main() {
	// Charger la configuration de l'application
	appConfig := config.LoadConfig()

	// Initialiser la base de données
	db := database.InitDB(appConfig.DatabaseURL)
	defer db.Close()

	// Initialiser la sécurité (JWT)
	jwtUtil := security.NewJWTUtil(appConfig.JWTSecret)

	// Créer une instance de l'application GoFiber
	app := routers.SetupApp(db, jwtUtil)

	// Démarrer le serveur sur le port spécifié
	err := app.Listen(":" + appConfig.Port)
	if err != nil {
		// Gérer les erreurs de démarrage du serveur
		panic(err)
	}
}
