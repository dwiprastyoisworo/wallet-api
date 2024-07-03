package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"wallet-api/migration"
	"wallet-api/src/infra/config"
)

func main() {

	conf := config.GetConfig()
	// Ganti dengan path ke file migrasi dan URL database Anda
	migrate := migration.NewMigration(conf)
	migrate.RunDatabaseMigration()
}
