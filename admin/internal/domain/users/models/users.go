package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UserName     string             `bson:"user_name" json:"user_name"`
	UserLastName string             `bson:"user_last_name" json:"user_last_name"`
	UserEmail    string             `bson:"user_email" json:"user_email"`
	Password     string             `bson:"password" json:"password,omitempty"`
	CreatedAt    string             `bson:"created_at" json:"created_at"`
	UpdatedAt    string             `bson:"updated_at" json:"updated_at"`
}

func (u *Users) UsersValidations(ctx *fiber.Ctx) error {

	if len(u.UserEmail) == 0 {
		return ctx.JSON(fiber.Map{"user_email": "User Email is required ..."})
	}
	if len(u.Password) < 6 {
		return ctx.JSON(fiber.Map{"password": "Password must be at least lenght of 6 elements ..."})
	}
	return nil
}
