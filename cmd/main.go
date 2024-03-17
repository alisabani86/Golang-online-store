package main

import (
	"log"
	"server/db"
	"server/internal/middleware"
	"server/internal/user"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database conection: %s", err)
	}

	middleware := middleware.NewMiddleware()
	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep, middleware)
	userHandler := user.NewHandler(userSvc, middleware)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")

}
