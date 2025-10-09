package services

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	if err := s.employeeDeleter.DeleteEmployee(ctx, id); err != nil {
		mapped := mapRepoErr(err)
		if mapped != nil && !errors.Is(mapped, err) {
			s.log.Error(LogMsgMappedError, zap.String(LogFieldOperation, OpDelete))
			return mapped
		}
		s.log.Error(LogMsgDeleteFailed, zap.Error(err))
		return err
	}
	return nil
}
