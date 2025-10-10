package http

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// GetEmployeeByDepartment godoc
// @Summary Get employees by company and department
// @Description Retrieves all employees for a specific company and department
// @Tags employees
// @Accept json
// @Produce json
// @Param companyID path int true "Company ID"
// @Param departmentID path int true "Department ID"
// @Success 200 {array} models.Employee "List of employees"
// @Failure 404 {object} map[string]string "Company or department not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /employees/company/{companyID}/department/{departmentID} [get]
func (h *Handler) getEmployeeByDepartment(c fiber.Ctx) error {
	companyID, err := strconv.Atoi(c.Params("companyID"))
	if err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}

	departmentID, err := strconv.Atoi(c.Params("departmentID"))
	if err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}

	employees, err := h.employeeService.GetEmployeeByDepartmentID(c.Context(), departmentID, companyID)
	if err != nil {
		return writeError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(employees)
}
