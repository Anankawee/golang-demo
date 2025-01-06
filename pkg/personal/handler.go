package personal

import "github.com/gofiber/fiber/v2"

func sendHeader(c *fiber.Ctx) error {
	return c.SendString(c.Method() + " Done")
}

func SetUpRoute(a *fiber.App) {
	option := a.Group("personal")
	option.Get("/list", sendHeader)
	option.Post("/add", sendHeader)
	option.Put("/modify", sendHeader)
	option.Patch("/update", sendHeader)
	option.Delete("/delete", sendHeader)
}
