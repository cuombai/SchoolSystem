package config

import (
	"context"
	"time"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
	MongoURI string
	DBName string
}

func LoadConfig() *AppConfig{
	godotenv.Load()//loads .env file
	return &AppConfig{
		Port: GetEnv("PORT", "8080"),
		MongoURI: GetEnv("MONGO_URI", ""),
		DBName: GetEnv("DB_NAME", "school"),
	}

}

func GetContext()(context.Context, context.CancelFunc){
	return context.WithTimeout(context.Background(), 10*time.Second)
}
