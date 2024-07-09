package main

import (
	"flag"
	"log"
	"note/internal/auth"
	"note/internal/infra/database"
	"note/internal/routes"
	"time"
)

const port = "8080"

type application struct {
	DSN          string
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	Domain       string
	APIKey       string
}

func main() {
	var app application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "verysecret", "signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain")
	flag.StringVar(&app.APIKey, "api-key", "", "api key")
	flag.Parse()

	database.ConnectDatabase()
	db := database.DB

	authConfig := &auth.Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Hour * 24,
		RefreshExpiry: time.Hour * 24 * 7,
		CookieDomain:  app.CookieDomain,
		CookiePath:    "/",
		CookieName:    "refresh_token",
	}

	router := routes.SetupRouter(db, authConfig)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
