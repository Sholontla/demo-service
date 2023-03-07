package c3

import (
	"service/producer/demo/monitoring/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func ProducerRouters3(app *fiber.App) {

	s := app.Group("service")
	p := s.Group("producer", middleware.RecordRequestLatency)
	app.Get("/swagger/*", swagger.HandlerDefault)

	p.Post("register3", SendDataServiceConsumer3)
}
