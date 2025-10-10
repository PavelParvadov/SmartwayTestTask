package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
)

func (s *EmployeeService) GetEmployeeByDepartmentID(ctx context.Context, departmentID, companyID int) ([]models.Employee, error) {
	employees, err := s.employeeProvider.GetEmployeeByDepartmentID(ctx, departmentID, companyID)
	if err != nil {
		mapped := mapRepoErr(err)
		if mapped != nil {
			s.log.Error("get employees by department error", zap.Error(mapped))
			return []models.Employee{}, mapped
		}

		s.log.Error("get employees by department error", zap.Error(err))
		return []models.Employee{}, err
	}
	return employees, nil
}
