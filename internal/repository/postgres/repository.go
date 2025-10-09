package postgres

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/dto"
	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
)

type EmployeeRepository interface {
	SaveEmployee(ctx context.Context, employee models.Employee) (int, error)
	UpdateEmployee(ctx context.Context, dto dto.UpdateEmployeeRequest) error
	DeleteEmployee(ctx context.Context, employeeId int) error
	GetEmployeesByCompanyID(ctx context.Context, companyId int) ([]models.Employee, error)
	GetEmployeeByDepartmentID(ctx context.Context, DepartmentId int, CompanyId int) ([]models.Employee, error)
}

type EmployeeRepositoryImpl struct {
	db *sql.DB
}

func NewEmployeeRepositoryImpl(db *sql.DB) *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{
		db: db,
	}
}
