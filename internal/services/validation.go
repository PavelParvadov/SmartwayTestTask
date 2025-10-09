package services

import (
	"fmt"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/dto"
	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
)

const (
	maxPhoneLen         = 20
	maxPassportFieldLen = 15
)

func validateEmployeeCreate(e models.Employee) error {
	if e.Name == "" {
		return fmt.Errorf("name is required")
	}
	if e.Surname == "" {
		return fmt.Errorf("surname is required")
	}
	if e.Phone == "" {
		return fmt.Errorf("phone is required")
	}
	if len(e.Phone) > maxPhoneLen {
		return fmt.Errorf("phone length must be <= %d", maxPhoneLen)
	}
	if e.CompanyId <= 0 {
		return fmt.Errorf("company_id must be > 0")
	}
	if e.DepartmentId <= 0 {
		return fmt.Errorf("department_id must be > 0")
	}
	if e.Passport.Type == "" || e.Passport.Number == "" {
		return fmt.Errorf("passport type and number are required")
	}
	if len(e.Passport.Type) > maxPassportFieldLen || len(e.Passport.Number) > maxPassportFieldLen {
		return fmt.Errorf("passport fields length must be <= %d", maxPassportFieldLen)
	}
	return nil
}

func validateEmployeeUpdate(req dto.UpdateEmployeeRequest) error {
	if req.Phone != "" && len(req.Phone) > maxPhoneLen {
		return fmt.Errorf("phone length must be <= %d", maxPhoneLen)
	}
	if req.PassportType != "" && len(req.PassportType) > maxPassportFieldLen {
		return fmt.Errorf("passport_type length must be <= %d", maxPassportFieldLen)
	}
	if req.PassportNumber != "" && len(req.PassportNumber) > maxPassportFieldLen {
		return fmt.Errorf("passport_number length must be <= %d", maxPassportFieldLen)
	}
	if req.DepartmentID < 0 {
		return fmt.Errorf("department_id must be >= 0")
	}
	if req.CompanyID < 0 {
		return fmt.Errorf("company_id must be >= 0")
	}
	return nil
}
