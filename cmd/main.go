package main

import (
	"context"
	"fmt"

	"github.com/canergulay/bilgipedia/internal/app"
	"github.com/canergulay/bilgipedia/internal/pkg/authentication"
	"github.com/canergulay/bilgipedia/internal/pkg/db"
	"github.com/canergulay/bilgipedia/internal/pkg/router"
	"github.com/canergulay/bilgipedia/internal/pkg/service"
	"github.com/canergulay/bilgipedia/internal/pkg/utils"
)

func main() {

	connectionString := utils.GetConnectionString("caner", "cKqxLQQfzOCRGAUo")
	DBName := utils.GetDBName()
	db, err := db.InitMongoDBConnection(connectionString, DBName)
	if err != nil {
		fmt.Printf("AN UNEXPECTED ERROR HAS OCCURED WHEN CONNECTING TO MONGODB %v", err)
	}

	auth := authentication.JwtManager{
		SecretKey: "bilgi",
	}

	sv := service.NewService(db, auth)

	router := router.InitializeRouter(context.Background(), sv, &auth)
	app.InitializeTheServer("8081", router.Routes, router.AuthMiddleware)

}
