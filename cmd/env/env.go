package env

import (
	"os"

	"github.com/dreamsparkx/go-web-boilerplate/internal/config"
	"github.com/joho/godotenv"
)

func LoadENV(app *config.AppConfig) {
	err := godotenv.Load(".env")
	if err != nil {
		godotenv.Load(".env.example")
	}
	if os.Getenv("APP_ENV") == "production" {
		(*app).InProduction = true
	} else {
		(*app).InProduction = false
	}
	if os.Getenv("USE_TEMPLATE_CACHE") == "true" {
		(*app).UseTemplateCache = true
	} else {
		(*app).UseTemplateCache = false
	}
	if os.Getenv("PORT") != "" {
		(*app).Port = os.Getenv("PORT")
	} else {
		(*app).Port = "8080"
	}
	(*app).DbCredentials.Host = os.Getenv("DB_HOST")
	(*app).DbCredentials.Name = os.Getenv("DB_NAME")
	(*app).DbCredentials.User = os.Getenv("DB_USER")
	(*app).DbCredentials.Password = os.Getenv("DB_PASSWORD")
	(*app).DbCredentials.SSLModeDisabled = os.Getenv("DB_SSL_MODE_DISABLED")
}
