package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/dto"
)

func (s *EmployeeService) UpdateEmployee(ctx context.Context, req dto.UpdateEmployeeRequest) error {
	if err := s.employeeUpdater.UpdateEmployee(ctx, req); err != nil {
		mapped := mapRepoErr(err)
		if mapped != nil {
			s.log.Error("update employee error", zap.Error(mapped))
			return mapped
		}

		s.log.Error("update employee error", zap.Error(err))
		return err
	}
	return nil
}
