package main

import (
	"log"
	"server/db"
	"server/internal/middleware"
	"server/internal/repository"
	"server/internal/service"

	router "server/router"
)

func main() {
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
	router.Start("0.0.0.0:8080")

}
