package api

import "github.com/gofiber/fiber/v2"

func AdminRouters(app *fiber.App) {

	s := app.Group("service")
	a := s.Group("admin")
	u := a.Group("user")

	u.Post("register", UserRegisterHandler)
	u.Get("get", GetUserHandler)
	u.Put("update", UpdateUserHandler)
	u.Delete("delete", DeleteUserHandler)

	u.Post("concurrent/register", ConcurrentUserRegisterHandler)
}
