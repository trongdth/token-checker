package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var config Config

func init() {
	err := godotenv.Load("./config/conf.yaml")
	if err != nil {
		godotenv.Load("../config/conf.yaml")
	}

	if err := env.Parse(&config); err != nil {
		log.Fatal("Error on parsing configuration file.", err)
	}

	log.Printf(`
		mongo_db_name: %s | mongo_db_url: %s
	`,
		config.MongoDBName, config.MongoDBURL,
	)
}

// GetConfig :
func GetConfig() *Config {
	return &config
}

// Config : struct
type Config struct {
	// mongo
	MongoDBURL  string `json:"mongo_db_url" env:"mongo_db_url"`
	MongoDBName string `json:"mongo_db_name" env:"mongo_db_name"`
}
