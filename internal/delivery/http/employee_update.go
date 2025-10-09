package http

import (
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/dto"
)

func (h *Handler) updateEmployee(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}
	var req dto.UpdateEmployeeRequest
	if err = c.Bind().Body(&req); err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}
	req.ID = id
	if err = h.employeeService.UpdateEmployee(c.Context(), req); err != nil {
		return writeError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
