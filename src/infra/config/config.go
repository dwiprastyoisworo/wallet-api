package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type PostgreDbConf struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
}

type Migration struct {
	Postgres bool
}

type Config struct {
	PostgreDb PostgreDbConf
	Migration Migration
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	requiredEnvVars := []string{
		"DB_PG_HOST",
		"DB_PG_PORT",
		"DB_PG_USERNAME",
		"DB_PG_PASSWORD",
		"DB_PG_DATABASE",
		"DB_PG_SSLMODE",
		"MIGRATION_POSTGRES",
	}

	envValues := make(map[string]string)

	for _, envVar := range requiredEnvVars {
		value := os.Getenv(envVar)
		if value == "" {
			panic(fmt.Sprintf("%s is not set", envVar))
		}
		envValues[envVar] = value
	}

	postgre := PostgreDbConf{
		Host:     envValues["DB_PG_HOST"],
		Port:     envValues["DB_PG_PORT"],
		User:     envValues["DB_PG_USERNAME"],
		Password: envValues["DB_PG_PASSWORD"],
		DbName:   envValues["DB_PG_DATABASE"],
		SslMode:  envValues["DB_PG_SSLMODE"],
	}

	migration := Migration{
		Postgres: envValues["MIGRATION_POSTGRES"] == "true",
	}

	return Config{
		PostgreDb: postgre,
		Migration: migration,
	}
}
