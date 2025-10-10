package http

import (
	"github.com/gofiber/fiber/v3"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
)

// CreateEmployee godoc
// @Summary Create a new employee
// @Description Creates a new employee with passport information
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body models.Employee true "Employee data"
// @Success 201 {object} map[string]int "Employee created successfully"
// @Failure 400 {object} map[string]string "Bad request - validation error"
// @Failure 409 {object} map[string]string "Conflict - employee with this phone or passport already exists"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /employees [post]
func (h *Handler) createEmployee(c fiber.Ctx) error {
	var req models.Employee
	if err := c.Bind().Body(&req); err != nil {
		return writeError(c, fiber.ErrBadRequest)
	}

	if !h.validateCreateEmployeeRequest(req) {
		return writeError(c, fiber.ErrBadRequest)
	}

	id, err := h.employeeService.SaveEmployee(c.Context(), req)
	if err != nil {
		return writeError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func (h *Handler) validateCreateEmployeeRequest(req models.Employee) bool {
	return req.Name != "" &&
		req.Surname != "" &&
		req.Phone != "" &&
		req.CompanyId > 0 &&
		req.DepartmentId > 0 &&
		req.Passport.Type != "" &&
		req.Passport.Number != "" &&
		len(req.Phone) <= 20 &&
		len(req.Passport.Type) <= 15 &&
		len(req.Passport.Number) <= 15
}
