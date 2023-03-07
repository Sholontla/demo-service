package api

import (
	"log"
	"service/admin/demo/internal/concurrent"
	"service/admin/demo/internal/domain/users/dao"
	"service/admin/demo/internal/domain/users/models"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func ConcurrentUserRegisterHandler(ctx *fiber.Ctx) error {

	var channel concurrent.StoreChannels
	channel.ChanUsers = make(chan models.Users, 10)

	var request models.Users

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.JSON(fiber.Map{"data": "Json not supported ..."})
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Request: ", request)
		dao.ConcurrentRegisterUser(request, channel.ChanUsers)
	}()

	var user []models.Users

	for users := range channel.ChanUsers {
		us := models.Users{
			ID:           users.ID,
			UserName:     users.UserName,
			UserLastName: users.UserLastName,
			UserEmail:    users.UserEmail,
			CreatedAt:    users.CreatedAt,
		}
		user = append(user, us)
	}
	response := map[string]interface{}{
		"user": user,
	}
	wg.Wait()
	return ctx.JSON(fiber.Map{"user": response})
}

func ConcurrentGetUserHandler(ctx *fiber.Ctx) error {

	var request map[string]string

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.JSON(fiber.Map{"data": "Json not supported ..."})
	}

	u := models.Users{
		UserEmail: request["user_email"],
	}
	us, err := dao.FindUserByEmail(u.UserEmail)
	if err != nil {
		return ctx.JSON(fiber.Map{"message": "Error while 'INSERTING' a user into 'RegisterUser' ...", "error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"user": us})
}

func ConcurrentUpdateUserHandler(ctx *fiber.Ctx) error {

	var request models.Users

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.JSON(fiber.Map{"data": "Json not supported ..."})
	}

	us, err := dao.UpdateUser(request, request.UserEmail)

	if err != nil {
		return ctx.JSON(fiber.Map{"message": "Error while 'INSERTING' a user into 'RegisterUser' ...", "error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"user": us})
}

func ConcurrentDeleteUserHandler(ctx *fiber.Ctx) error {

	var request models.Users

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.JSON(fiber.Map{"data": "Json not supported ..."})
	}

	us, err := dao.DeleteUser(request.UserEmail)

	if err != nil {
		return ctx.JSON(fiber.Map{"message": "Error while 'INSERTING' a user into 'RegisterUser' ...", "error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"user": us})
}
