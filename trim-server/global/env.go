package global

import (
	"os"

	"github.com/joho/godotenv"
)

func Env(env string) string {

	godotenv.Load("./.env")

	envVars := map[string]string{
		"ADDRESS":    os.Getenv("ADDRESS"),
		"JWT_SECRET": os.Getenv("JWT_Secret"),
		"DB_URL":     os.Getenv("DB_URL"),
	}

	return envVars[env]

}
