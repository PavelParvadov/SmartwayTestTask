package services

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
)

func (s *EmployeeService) GetEmployeesByCompanyID(ctx context.Context, companyID int) ([]models.Employee, error) {
	employees, err := s.employeeProvider.GetEmployeesByCompanyID(ctx, companyID)
	if err != nil {
		mapped := mapRepoErr(err)
		if mapped != nil && !errors.Is(mapped, err) {
			s.log.Error(LogMsgMappedError, zap.String(LogFieldOperation, OpGetByCompany))
			return nil, mapped
		}
		s.log.Error(LogMsgGetByCompanyFailed, zap.Error(err))
		return nil, err
	}
	return employees, nil
}
