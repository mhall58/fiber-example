package input

import (
	"github.com/gofiber/fiber/v2"
	"itemsapi/pkg/input"
)

type InputHandler struct {
	Repo input.Repository
}

func (h *InputHandler) List(c *fiber.Ctx) error {

	inputs, err := h.Repo.GetAll(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": inputs,
	})
}
