package main

import (
	"log"

	"github.com/JhonMos/alumnitweet/bd"
	"github.com/JhonMos/alumnitweet/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
