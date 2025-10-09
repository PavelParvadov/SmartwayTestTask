package http

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) getEmployeesByCompany(c fiber.Ctx) error {
	companyIDStr := c.Params("companyID")
	companyID, err := strconv.Atoi(companyIDStr)
	if err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}
	employees, err := h.employeeService.GetEmployeesByCompanyID(c.Context(), companyID)
	if err != nil {
		return writeError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(employees)
}
