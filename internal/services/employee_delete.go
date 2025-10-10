package services

import (
	"context"

	"go.uber.org/zap"
)

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	if err := s.employeeDeleter.DeleteEmployee(ctx, id); err != nil {
		mapped := mapRepoErr(err)
		if mapped != nil {
			s.log.Error("delete employee error", zap.Error(mapped))
			return mapped
		}

		s.log.Error("delete employee error", zap.Error(err))
		return err
	}
	return nil
}
