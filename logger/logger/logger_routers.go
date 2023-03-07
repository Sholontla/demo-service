package logger

import "github.com/gofiber/fiber/v2"

func SupplierRouters(r *fiber.App) {

	re := DBRequestService{}

	api := r.Group("api")
	logger := api.Group("logger")

	logger.Post("create", re.LoggerCreateHandler)
	logger.Get("get", re.LoggerGetHandler)
	logger.Get("get/all", re.LoggerGetAllHandler)

}
