package main

import (
	"auth_microservice/internal/repository"
	"auth_microservice/internal/service"
	"auth_microservice/internal/transport"
	"auth_microservice/pkg/database"
	"auth_microservice/pkg/hash"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

// to be moved to environment variables
var (
	salt       = "salt"
	hmacSecret = "secret"
	accessTTL  = 15 * time.Minute
	refreshTTL = 720 * time.Hour
)

func main() {
	db, err := database.NewPostgresConnection(database.ConnectionInfo{Host: "localhost", Port: 5432, UserName: "crud-6", DBName: "crud-6-db", SSLMode: "disable", Password: "12345"})
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	hasher := hash.NewSHA1Hasher(salt)

	usersRepo := repository.NewUsers(db)
	tokensRepo := repository.NewTokens(db)
	usersService := service.NewUsers(usersRepo, tokensRepo, hasher, []byte(hmacSecret), accessTTL, refreshTTL)
	handler := transport.NewHandler(usersService)

	server := &http.Server{Addr: "localhost:8080", Handler: handler.InitRoutes()}
	server.ListenAndServe()
}
