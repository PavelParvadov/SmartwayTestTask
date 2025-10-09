package postgres

import (
	"context"
	"errors"

	"github.com/lib/pq"
	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
	"github.com/pavelParvadov/SmartwayTask/internal/repository/postgres/scripts"
)

func (r *EmployeeRepositoryImpl) SaveEmployee(ctx context.Context, employee models.Employee) (int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var exists bool
	// проверяем, сущствует ли компания по переданному company_id
	if err = tx.QueryRowContext(ctx, scripts.IsCompanyExists, employee.CompanyId).Scan(&exists); err != nil {
		return 0, err
	}
	if !exists {
		return 0, ErrCompanyNotFound
	}

	// проверяем уникальность паспорта по (type, number)
	if err = tx.QueryRowContext(ctx, scripts.IsPassportExists, employee.Passport.Type, employee.Passport.Number).Scan(&exists); err != nil {
		return 0, err
	}
	if exists {
		return 0, ErrPassportExist
	}

	var passportID int
	// создаём паспорт и получаем его id
	if err = tx.QueryRowContext(ctx, scripts.InsertPassport, employee.Passport.Type, employee.Passport.Number).Scan(&passportID); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && string(pqErr.Code) == "23505" {
			return 0, ErrPassportExist
		}
		return 0, err
	}

	var depExists bool
	// проверяем, существуе ли отдел по переданному department_id
	if err = tx.QueryRowContext(ctx, scripts.IsDepartmentExists, employee.DepartmentId).Scan(&depExists); err != nil {
		return 0, err
	}
	if !depExists {
		return 0, ErrDepartmentNotFound
	}
	departmentID := employee.DepartmentId

	var employeeID int
	// создаём сотрудника
	if err = tx.QueryRowContext(ctx, scripts.InsertEmployee,
		employee.Name,
		employee.Surname,
		employee.Phone,
		employee.CompanyId,
		passportID,
		departmentID,
	).Scan(&employeeID); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && string(pqErr.Code) == "23505" {
			return 0, ErrPhoneExist
		}
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return employeeID, nil
}
