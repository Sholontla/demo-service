package api

import (
	"fmt"
	dto "service/files_server/demo/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func PipDataHandler(ctx *fiber.Ctx) error {

	var request dto.PipeServerModel

	if err := ctx.BodyParser(&request); err != nil {
		ctx.JSON(fiber.Map{"error_reques": fmt.Sprintf("Error JSON request: %v\n", err)})
	}

	return ctx.JSON(fiber.Map{"produce_message": request})
}
