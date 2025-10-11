package http

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// GetEmployeesByCompany godoc
// @Summary Get employees by company
// @Description Retrieves all employees for a specific company
// @Tags employees
// @Accept json
// @Produce json
// @Param companyID path int true "Company ID"
// @Success 200 {array} models.Employee "List of employees"
// @Failure 404 {string} string "Company not found"
// @Failure 500 {string} string "Internal server error"
// @Router /employees/company/{companyID} [get]
func (h *Handler) getEmployeesByCompany(c fiber.Ctx) error {
	companyID, err := strconv.Atoi(c.Params("companyID"))
	if err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}

	employees, err := h.employeeService.GetEmployeesByCompanyID(c.Context(), companyID)
	if err != nil {
		return writeError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(employees)
}
