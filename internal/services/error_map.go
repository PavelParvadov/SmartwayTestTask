package services

import (
	"errors"
	"fmt"

	repo "github.com/pavelParvadov/SmartwayTask/internal/repository/postgres"
)

func mapRepoErr(err error) error {
	switch {
	case errors.Is(err, repo.ErrCompanyNotFound):
		return ErrSvcInvalidInput
	case errors.Is(err, repo.ErrDepartmentNotFound):
		return ErrSvcInvalidInput
	case errors.Is(err, repo.ErrEmployeeNotFound):
		return ErrSvcEmployeeNotFound
	case errors.Is(err, repo.ErrPassportExist), errors.Is(err, repo.ErrPhoneExist):
		return fmt.Errorf("%w: %w", ErrSvcConflict, err)
	default:
		return err
	}
}
