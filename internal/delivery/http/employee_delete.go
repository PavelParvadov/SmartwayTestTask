package http

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Deletes an employee by ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 204 "Employee deleted successfully"
// @Failure 404 {string} string "Employee not found"
// @Failure 500 {string} string "Internal server error"
// @Router /employees/{id} [delete]
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
