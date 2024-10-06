package main

import (
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"

	"log"
)

func main() {
	dbCon, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Не инициализирволась БД")
	}

	userRep := user.NewRepository(dbCon.GetDB())
	userSrv := user.NewService(userRep)
	userHandler := user.NewHandler(userSrv)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8081")

}
