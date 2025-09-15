package config

type AppConfig struct {
	Port string
	MongoURI string
	DBName string
}

func LoadConfig() *AppConfig{
	return &AppConfig{
		Port: GetEnv("PORT", "8080"),
		MongoURI: GetEnv("MONGO_URI", ""),
		DBName: GetEnv("DB_NAME", "school"),
	}

}