package main

import (
	"github.com/YurcheuskiRadzivon/HSC-pattern/internal/controller"
	"github.com/YurcheuskiRadzivon/HSC-pattern/internal/handler"
	"github.com/YurcheuskiRadzivon/HSC-pattern/internal/repository"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	userRepo, err := repository.NewUserRepository("asd", "", "", "", "")
	if err != nil {
		println(err)
	}
	userController := controller.NewUserController(userRepo)
	userHandler := handler.NewUserHandler(userController)
	htmlengine := html.New("../web/templates", ".html")
	app := fiber.New(fiber.Config{
		Views: htmlengine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home_p", nil)
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login_p", nil)
	})
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register_p", nil)
	})
	app.Post("/register", userHandler.InsertUser)

	filepath := filepath.Join("../web/static")
	app.Static("/", filepath)
	app.Listen(":8080")
}
