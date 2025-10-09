package http

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) deleteEmployee(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}
	if err = h.employeeService.DeleteEmployee(c.Context(), id); err != nil {
		return writeError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

