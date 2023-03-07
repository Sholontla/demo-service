package server

import (
	"log"
	"os"
	"os/signal"
	"service/producer/demo/api/c1"
	"service/producer/demo/api/c2"
	"service/producer/demo/api/c3"
	"service/producer/demo/api/c4"
	"service/producer/demo/api/c5"
	"service/producer/demo/monitoring"
	"service/producer/demo/monitoring/middleware"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func ProducerServerStart() {

	app := fiber.New()

	c1.ProducerRouters1(app)
	c2.ProducerRouters2(app)
	c3.ProducerRouters3(app)
	c4.ProducerRouters4(app)
	c5.ProducerRouters5(app)
	monitoring.PrometheusRoute(app)
	middleware.RegisterPrometheusMetrics()

	go app.Listen(":8004")

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
