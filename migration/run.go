package migration

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"wallet-api/src/infra/config"
)

type migration struct {
	config config.Config
}

func NewMigration(config config.Config) migration {
	return migration{config: config}
}

func (m migration) RunDatabaseMigration() {
	if m.config.Migration.Postgres {
		postgreMigration(m.config.PostgreDb)
	}
}

func postgreMigration(psConfig config.PostgreDbConf) {
	// Menyusun URL koneksi PostgreSQL dari parameter config
	postgresURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		psConfig.User,
		psConfig.Password,
		psConfig.Host,
		psConfig.Port,
		psConfig.DbName,
		psConfig.SslMode,
	)

	m, err := migrate.New(
		"file://migration/postgre",
		postgresURL,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Menjalankan migrasi ke atas
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migration needed")
		} else {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
	} else {
		log.Println("Migrations applied successfully")
	}

	// Jika ingin menjalankan migrasi ke bawah
	/*
	   if err := m.Down(); err != nil {
	       log.Fatalf("Failed to rollback migrations: %v", err)
	   } else {
	       log.Println("Migrations rolled back successfully")
	   }
	*/
}
