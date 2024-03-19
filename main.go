package main

import (
	"log"
	"online_store/db"
	"online_store/internal/middleware"
	"online_store/internal/repository"
	"online_store/internal/service"
	"os"

	router "online_store/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database conection: %s", err)
	}

	middleware := middleware.NewMiddleware()

	Svc := service.NewService(
		repository.NewRepository(dbConn.GetDB()),
		middleware,
	)
	Handler := router.NewHandler(Svc, middleware)

	router.InitRouter(Handler)
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")

	router.Start(HOST + ":" + PORT)

}
