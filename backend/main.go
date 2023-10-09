package main

import (
	"github.com/Akito-Fujihara/ts-go-sandbox/database"
	"github.com/Akito-Fujihara/ts-go-sandbox/router"
)

func main() {
	db, _ := database.DB.DB()
	defer db.Close()

	e := router.NewRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
