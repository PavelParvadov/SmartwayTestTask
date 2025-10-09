package services

import (
	"context"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/dto"
	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
	"go.uber.org/zap"
)

type EmployeeService struct {
	log              *zap.Logger
	employeeSaver    EmployeeSaver
	employeeProvider EmployeeProvider
	employeeDeleter  EmployeeDeleter
	employeeUpdater  EmployeeUpdater
}

func NewEmployeeService(
	log *zap.Logger,
	employeeSaver EmployeeSaver,
	employeeProvider EmployeeProvider,
	employeeDeleter EmployeeDeleter,
	employeeUpdater EmployeeUpdater) *EmployeeService {

	return &EmployeeService{
		log:              log,
		employeeSaver:    employeeSaver,
		employeeProvider: employeeProvider,
		employeeDeleter:  employeeDeleter,
		employeeUpdater:  employeeUpdater,
	}
}

type EmployeeSaver interface {
	SaveEmployee(ctx context.Context, employee models.Employee) (int, error)
}
type EmployeeProvider interface {
	GetEmployeesByCompanyID(ctx context.Context, companyID int) ([]models.Employee, error)
	GetEmployeeByDepartmentID(ctx context.Context, departmentID, companyID int) ([]models.Employee, error)
}

type EmployeeDeleter interface {
	DeleteEmployee(ctx context.Context, id int) error
}

type EmployeeUpdater interface {
	UpdateEmployee(ctx context.Context, employee dto.UpdateEmployeeRequest) error
}
