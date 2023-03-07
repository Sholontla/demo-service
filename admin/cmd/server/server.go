package server

import (
	"log"
	"os"
	"os/signal"
	"service/admin/demo/api"
	"service/admin/demo/internal/db"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func ProducerServerStart() {
	if db.CheckConection() == 0 {
		log.Fatal("No DB Connection ...")
		return
	}

	app := fiber.New()

	api.AdminRouters(app)

	go app.Listen(":8005")

	// SIGINT is the signal sent when we press Ctrl+C
	// SIGTERM gracefully kills the process
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down Service Demo server.....")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Shutting Down Service Demo Server: %v\n", err)
	}
}
