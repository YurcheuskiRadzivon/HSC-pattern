package handler

import (
	"context"
	"github.com/YurcheuskiRadzivon/HSC-pattern/internal/controller"
	"github.com/YurcheuskiRadzivon/HSC-pattern/model"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUser(c *fiber.Ctx) error
	InsertUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetUserPassword(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
}
type userHandler struct {
	ctx        context.Context
	controller controller.UserController
}

func NewUserHandler(controller controller.UserController) UserHandler {
	return &userHandler{
		controller: controller,
		ctx:        context.Background(),
	}

}
func (us *userHandler) GetUser(c *fiber.Ctx) error {
	return nil
}
func (us *userHandler) InsertUser(c *fiber.Ctx) error {
	var User model.User
	if err := c.BodyParser(&User); err != nil {
		return c.SendString("Error ")
	}

}
func (us *userHandler) UpdateUser(c *fiber.Ctx) error {
	return nil
}
func (us *userHandler) DeleteUser(c *fiber.Ctx) error {
	return nil
}
func (us *userHandler) GetUserPassword(c *fiber.Ctx) error {
	return nil
}
func (us *userHandler) LoginUser(c *fiber.Ctx) error {
	return nil
}
