package services

import "errors"

var (
	ErrSvcEmployeeNotFound   = errors.New("employee not found")
	ErrSvcCompanyNotFound    = errors.New("company not found")
	ErrSvcDepartmentNotFound = errors.New("department not found")
	ErrSvcConflict           = errors.New("conflict")
	ErrSvcInvalidInput       = errors.New("invalid input")
)
