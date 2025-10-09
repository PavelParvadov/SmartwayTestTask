package http

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) getEmployeeByDepartment(c fiber.Ctx) error {
	companyIDStr := c.Params("companyID")
	departmentIDStr := c.Params("departmentID")
	companyID, err := strconv.Atoi(companyIDStr)
	if err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}
	departmentID, err := strconv.Atoi(departmentIDStr)
	if err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}
	employees, err := h.employeeService.GetEmployeeByDepartmentID(c.Context(), departmentID, companyID)
	if err != nil {
		return writeError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(employees)
}
