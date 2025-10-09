package postgres

import "errors"

const (
	UniqueViolationCode = "23505"
)

var (
	ErrPassportExist      = errors.New("passport already exists")
	ErrPhoneExist         = errors.New("phone already exists")
	ErrCompanyNotFound    = errors.New("company not found")
	ErrDepartmentNotFound = errors.New("department not found")
	ErrEmployeeNotFound   = errors.New("employee not found")
)
