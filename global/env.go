package global

import (
	"os"

	"github.com/joho/godotenv"
)

func Env(env string) string {

	godotenv.Load("./.env")

	envVars := map[string]string{
		"PORT":       os.Getenv("PORT"),
		"JWT_SECRET": os.Getenv("JWT_SECRET"),
		"DB_URL":     os.Getenv("DB_URL"),
	}

	return envVars[env]

}
