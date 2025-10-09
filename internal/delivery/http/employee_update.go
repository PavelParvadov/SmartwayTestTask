package http

import (
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/dto"
)

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
