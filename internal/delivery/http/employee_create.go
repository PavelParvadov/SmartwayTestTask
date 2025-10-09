package http

import (
	"github.com/gofiber/fiber/v3"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
)

func (h *Handler) createEmployee(c fiber.Ctx) error {
	var req models.Employee
	if err := c.Bind().Body(&req); err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}
	id, err := h.employeeService.SaveEmployee(c.Context(), req)
	if err != nil {
		return writeError(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}
