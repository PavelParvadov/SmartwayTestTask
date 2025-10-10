package http

import (
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/dto"
)

// UpdateEmployee godoc
// @Summary Update an employee
// @Description Updates an existing employee
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body dto.UpdateEmployeeRequest true "Employee update data"
// @Success 204 "Employee updated successfully"
// @Failure 400 {object} map[string]string "Bad request - validation error"
// @Failure 404 {object} map[string]string "Employee not found"
// @Failure 409 {object} map[string]string "Conflict - employee with this phone or passport already exists"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /employees/{id} [patch]
func (h *Handler) updateEmployee(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}

	var req dto.UpdateEmployeeRequest
	if err = c.Bind().Body(&req); err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}

	if !h.validateUpdateEmployeeRequest(req) {
		return writeError(c, fiber.ErrBadRequest)
	}

	req.ID = id
	if err = h.employeeService.UpdateEmployee(c.Context(), req); err != nil {
		return writeError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) validateUpdateEmployeeRequest(req dto.UpdateEmployeeRequest) bool {
	if req.Phone != "" && len(req.Phone) > 20 {
		return false
	}
	if req.PassportType != "" && len(req.PassportType) > 15 {
		return false
	}
	if req.PassportNumber != "" && len(req.PassportNumber) > 15 {
		return false
	}
	if req.DepartmentID < 0 || req.CompanyID < 0 {
		return false
	}
	return true
}
