package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
)

func (s *EmployeeService) SaveEmployee(ctx context.Context, employee models.Employee) (int, error) {
	id, err := s.employeeSaver.SaveEmployee(ctx, employee)
	if err != nil {
		mapped := mapRepoErr(err)
		if mapped != nil {
			s.log.Error("save employee error", zap.Error(mapped))
			return 0, mapped
		}

		s.log.Error("save employee error", zap.Error(err))
		return 0, err
	}
	return id, nil
}
