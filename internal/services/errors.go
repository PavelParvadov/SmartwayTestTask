package services

import (
	"errors"
	repo "github.com/pavelParvadov/SmartwayTask/internal/repository/postgres"
)

var (
	ErrSvcEmployeeNotFound   = errors.New("employee not found")
	ErrSvcCompanyNotFound    = errors.New("company not found")
	ErrSvcDepartmentNotFound = errors.New("department not found")
	ErrSvcPassportExist      = errors.New("passport already exists")
	ErrSvcPhoneExist         = errors.New("phone already exists")
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
