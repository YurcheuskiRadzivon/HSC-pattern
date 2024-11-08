package routes

import (
	"github.com/YurcheuskiRadzivon/HSC-pattern/api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"path/filepath"
)

func NewFiberRouter(userHandler handler.UserHandler, propertyHandler handler.PropertyHandler) *fiber.App {
	htmlengine := html.New("../../web/templates", ".html")
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
	app.Post("/login", userHandler.LoginUser)

	filepath := filepath.Join("..", "..", "web", "static")
	app.Static("/", filepath)
	return app
}
