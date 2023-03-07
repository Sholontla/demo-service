package logger

import (
	"strconv"

	utilerrors "service/logger/service/utils/util_errors"

	"github.com/gofiber/fiber/v2"
)

type DBRequestService struct {
	data *Config
}

var request LogModel

func (e DBRequestService) LoggerCreateHandler(ctx *fiber.Ctx) error {

	if err := ctx.BodyParser(&request); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	service := e.data.LoggerModel.LogEntry.Insert(request.LogEntry)

	return ctx.JSON(fiber.Map{"message": service})
}

func (e DBRequestService) LoggerGetHandler(ctx *fiber.Ctx) error {

	if err := ctx.BodyParser(&request.LogEntry.ID); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	service, err := e.data.LoggerModel.LogEntry.GetOne(request.LogEntry.ID)
	if err != nil {
		jsonErr := utilerrors.NewInternalServerError("Invalid ID", err)
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	return ctx.JSON(fiber.Map{"message": service})
}

func (e DBRequestService) LoggerGetAllHandler(ctx *fiber.Ctx) error {

	service, err := e.data.LoggerModel.LogEntry.All()
	if err != nil {
		jsonErr := utilerrors.NewInternalServerError("Invalid ID", err)
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	return ctx.JSON(fiber.Map{"message": service})
}

func (e DBRequestService) LoggerUpdateInfoHandler(ctx *fiber.Ctx) error {
	if err := ctx.BodyParser(&request.LogEntry.ID); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	service, err := e.data.LoggerModel.LogEntry.UpdateCollection()
	if err != nil {
		jsonErr := utilerrors.NewInternalServerError("Invalid ID", err)
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	return ctx.JSON(fiber.Map{"message": service})
}
