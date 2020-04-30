package main

import (
	"log"

	"github.com/kdmc-lean-tech/eevee-go/db"
	"github.com/kdmc-lean-tech/eevee-go/handlers"
)

func main() {
	if db.CheckConnection() == false {
		log.Fatal("No connection to the database")
		return
	}
	handlers.RouteManagement()
}
