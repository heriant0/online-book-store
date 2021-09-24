package main

import (
	"log"
	"online-book-store/config"
)

func main() {
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal(err.Error())
	}

	r:= config.CreateRouter()

	config.InitRouter(db,r).InitializeRoutes()

	config.Run(r)

}
