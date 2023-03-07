package server

import (
	"log"
	"os"
	"os/signal"
	"service/files_server/demo/api"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func PipeFileServerStart() {

	app := fiber.New()

	api.PipeFilesRouters(app)
	api.PipeDataRouters(app)
	go app.Listen(":8006")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down Service Pipe Files Server.....")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Shutting Down Service Pipe Files Server: %v\n", err)
	}
}
