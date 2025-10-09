package services

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
)

func (s *EmployeeService) GetEmployeeByDepartmentID(ctx context.Context, departmentID, companyID int) ([]models.Employee, error) {
	employees, err := s.employeeProvider.GetEmployeeByDepartmentID(ctx, departmentID, companyID)
	if err != nil {
		mapped := mapRepoErr(err)
		if mapped != nil && !errors.Is(mapped, err) {
			s.log.Error(LogMsgMappedError, zap.String(LogFieldOperation, OpGetByDepartment))
			return nil, mapped
		}
		s.log.Error(LogMsgGetByDepartmentFailed, zap.Error(err))
		return nil, err
	}
	return employees, nil
}
