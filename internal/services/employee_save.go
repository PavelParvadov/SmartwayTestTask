package services

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
)

func (s *EmployeeService) SaveEmployee(ctx context.Context, employee models.Employee) (int, error) {
	if err := validateEmployeeCreate(employee); err != nil {
		s.log.Error(LogMsgInvalidInput)
		return 0, ErrSvcInvalidInput
	}
	id, err := s.employeeSaver.SaveEmployee(ctx, employee)
	if err != nil {
		mapped := mapRepoErr(err)
		if mapped != nil && !errors.Is(mapped, err) {
			s.log.Error(LogMsgMappedError, zap.String(LogFieldOperation, OpSave))
			return 0, mapped
		}
		s.log.Error(LogMsgSaveFailed, zap.Error(err))
		return 0, err
	}
	return id, nil
}
