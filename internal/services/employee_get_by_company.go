package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
)

func (s *EmployeeService) GetEmployeesByCompanyID(ctx context.Context, companyID int) ([]models.Employee, error) {
	employees, err := s.employeeProvider.GetEmployeesByCompanyID(ctx, companyID)
	if err != nil {
		mapped := mapRepoErr(err)
		if mapped != nil {
			s.log.Error("get employees by company error", zap.Error(mapped))
			return []models.Employee{}, mapped
		}

		s.log.Error("get employees by company error", zap.Error(err))
		return []models.Employee{}, err
	}
	return employees, nil
}
