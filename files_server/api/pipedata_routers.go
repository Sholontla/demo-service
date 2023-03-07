package api

import "github.com/gofiber/fiber/v2"

func PipeDataRouters(app *fiber.App) {

	api := app.Group("api")
	pipR := api.Group("pipedata")

	pipR.Post("register", PipDataHandler)

}
