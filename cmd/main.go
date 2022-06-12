package main

import (
	"github.com/egorrridze/payment-emulator"
	"github.com/egorrridze/payment-emulator/pkg/handler"
	"log"
)

func main() {
	handlers :=  new(handler.Handler)
	srv := new(emulator.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}