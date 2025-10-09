package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/dto"
)

func (s *EmployeeService) UpdateEmployee(ctx context.Context, req dto.UpdateEmployeeRequest) error {
	if err := validateEmployeeUpdate(req); err != nil {
		s.log.Error(LogMsgInvalidInput)
		return ErrSvcInvalidInput
	}
	if err := s.employeeUpdater.UpdateEmployee(ctx, req); err != nil {
		mapped := mapRepoErr(err)
		if mapped != nil {
			s.log.Error(LogMsgMappedError, zap.String(LogFieldOperation, OpUpdate))
			return mapped
		}
		s.log.Error(LogMsgUpdateFailed, zap.Error(err))
		return err
	}
	return nil
}
