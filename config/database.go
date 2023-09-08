package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// Remplacez les valeurs suivantes par les informations de connexion à votre base de données PostgreSQL
	dsn := "host=your-database-host user=your-database-user password=your-database-password dbname=your-database-name port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Impossible de se connecter à la base de données : " + err.Error())
	}

	DB = db
}
