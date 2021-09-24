package main

import (
	"log"
	"online-book-store/config"
)

func main() {
	_, err := config.ConnectDB()

	if err != nil {
		log.Fatal(err.Error())
	}

	

}
