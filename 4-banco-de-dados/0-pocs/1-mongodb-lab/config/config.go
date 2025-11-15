package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBHost             string `env:"MONGO_HOST,required"`
	DBPort             string `env:"MONGO_PORT,required"`
	DBUser             string `env:"MONGO_USER,required"`
	DBPass             string `env:"MONGO_PASS,required"`
	DatabaseName       string `env:"DATABASE_NAME,required"`
	UserCollectionName string `env:"USER_COLLECTION_NAME,required"`
}

func LoadConfig() (*AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, utilizando variáveis de ambiente do sistema")
	}

	opts := env.Options{
		RequiredIfNoDef: true,
	}

	config := &AppConfig{}

	if err := env.ParseWithOptions(config, opts); err != nil {
		return nil, fmt.Errorf("Erro ao carregar configuração: %v", err)
	}

	fmt.Println("Config loaded successfully")
	fmt.Println("Host:", config.DBHost)
	fmt.Println("Port:", config.DBPort)
	fmt.Println("User:", config.DBUser)

	return config, nil
}
