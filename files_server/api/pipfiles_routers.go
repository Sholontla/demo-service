package api

import "github.com/gofiber/fiber/v2"

func PipeFilesRouters(app *fiber.App) {

	api := app.Group("api")
	pipR := api.Group("pipefile")

	pipR.Post("register", PipeFileHandler)

}
