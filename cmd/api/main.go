package main

import (
	"github.com/gofiber/fiber/v2"
	"itemsapi/cmd/api/handlers/input"
	input2 "itemsapi/pkg/input"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//bio - gas
	bioGas := app.Group("/biogas")
	inputsHandler := &input.InputHandler{
		Repo: &input2.StaticInputs{},
	}
	bioGas.Get("/inputs", inputsHandler.List)

	app.Listen(":3000")
}
