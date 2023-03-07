package api

import (
	"log"
	"service/admin/demo/internal/concurrent"
	"service/admin/demo/internal/domain/users/dao"
	"service/admin/demo/internal/domain/users/models"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func UserRegisterHandler(ctx *fiber.Ctx) error {

	var channel concurrent.StoreChannels
	channel.ChanUsers = make(chan models.Users, 10)

	var request models.Users

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.JSON(fiber.Map{"data": "Json not supported ..."})
	}

	if len(request.UserEmail) == 0 {
		return ctx.JSON(fiber.Map{"user_email": "User Email is required ..."})
	}
	if len(request.Password) < 6 {
		return ctx.JSON(fiber.Map{"password": "Password must be at least lenght of 6 elements ..."})
	}

	_, exists, _ := dao.ExistsUser(request.UserEmail)
	if exists == false {
		return ctx.JSON(fiber.Map{"user": "User Exists, choose another one ..."})
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Request: ", request)
		dao.ConcurrentRegisterUser(request, channel.ChanUsers)
	}()

	response, register, err := dao.RegisterUser(request)
	if err != nil {
		return ctx.JSON(fiber.Map{"message": "Error while 'INSERTING' a user into 'RegisterUser' ...", "error": err.Error()})
	}
	if register == false {
		return ctx.JSON(fiber.Map{"error": "Error while 'INSERTING' a user into 'RegisterUser' ..."})
	}

	return ctx.JSON(fiber.Map{"user": response})
}

func GetUserHandler(ctx *fiber.Ctx) error {

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

func UpdateUserHandler(ctx *fiber.Ctx) error {

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

func DeleteUserHandler(ctx *fiber.Ctx) error {

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
