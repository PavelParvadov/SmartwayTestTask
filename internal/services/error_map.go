package services

import (
	"errors"

	repo "github.com/pavelParvadov/SmartwayTask/internal/repository/postgres"
)

func mapRepoErr(err error) error {
	switch {
	case errors.Is(err, repo.ErrCompanyNotFound):
		return ErrSvcCompanyNotFound
	case errors.Is(err, repo.ErrDepartmentNotFound):
		return ErrSvcDepartmentNotFound
	case errors.Is(err, repo.ErrEmployeeNotFound):
		return ErrSvcEmployeeNotFound
	case errors.Is(err, repo.ErrPassportExist):
		return ErrSvcPassportExist
	case errors.Is(err, repo.ErrPhoneExist):
		return ErrSvcPhoneExist
	default:
		return err
	}
}
