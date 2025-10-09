package http

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/pavelParvadov/SmartwayTask/internal/domain/dto"
	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
)

type Handler struct {
	employeeService EmployeeService
}

func NewHandler(employeeService EmployeeService) *Handler {
	return &Handler{employeeService: employeeService}
}

type EmployeeService interface {
	SaveEmployee(ctx context.Context, employee models.Employee) (int, error)
	GetEmployeesByCompanyID(ctx context.Context, companyID int) ([]models.Employee, error)
	GetEmployeeByDepartmentID(ctx context.Context, departmentID, companyID int) ([]models.Employee, error)
	DeleteEmployee(ctx context.Context, id int) error
	UpdateEmployee(ctx context.Context, employee dto.UpdateEmployeeRequest) error
}

func (h *Handler) Register(api fiber.Router) {
	g := api.Group("/employees")
	g.Post("", h.createEmployee)
	g.Get("/company/:companyID/department/:departmentID", h.getEmployeeByDepartment)
	g.Get("/company/:companyID", h.getEmployeesByCompany)
	g.Patch("/:id", h.updateEmployee)
	g.Delete("/:id", h.deleteEmployee)
}
