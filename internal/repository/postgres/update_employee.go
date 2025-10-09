package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/pavelParvadov/SmartwayTask/internal/domain/dto"
	"github.com/pavelParvadov/SmartwayTask/internal/repository/postgres/scripts"
)

func (r *EmployeeRepositoryImpl) UpdateEmployee(ctx context.Context, req dto.UpdateEmployeeRequest) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var passportID int
	// получаем текущий passport_id сотрудника
	if err = tx.QueryRowContext(ctx, scripts.QueryFindPassportID, req.ID).Scan(&passportID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrEmployeeNotFound
		}
		return err
	}

	var (
		name          string
		surname       string
		phone         string
		companyID     int
		curPassType   string
		curPassNumber string
		depName       string
		depPhone      string
		departmentID  int
	)

	// получаем данные сотрудника на текщий момент
	if err = tx.QueryRowContext(ctx, scripts.QueryFindEmployee, req.ID).Scan(
		&name,
		&surname,
		&phone,
		&companyID,
		&curPassType,
		&curPassNumber,
		&depName,
		&depPhone,
		&departmentID,
	); err != nil {
		return err
	}

	if req.Name != "" {
		name = req.Name
	}
	if req.Surname != "" {
		surname = req.Surname
	}
	if req.Phone != "" {
		phone = req.Phone
	}
	if req.CompanyID != 0 {
		var exists bool
		// проверяем, существует ли компания по переданному company_id
		if err = tx.QueryRowContext(ctx, scripts.IsCompanyExists, req.CompanyID).Scan(&exists); err != nil {
			return err
		}
		if !exists {
			return ErrCompanyNotFound
		}
		companyID = req.CompanyID
	}
	if req.DepartmentID != 0 {
		var exists bool
		// проверяем, существует ли отдел по переданному id департамента
		if err = tx.QueryRowContext(ctx, scripts.IsDepartmentExists, req.DepartmentID).Scan(&exists); err != nil {
			return err
		}
		if !exists {
			return ErrDepartmentNotFound
		}
		departmentID = req.DepartmentID
	}

	if req.PassportType != "" || req.PassportNumber != "" {
		newType := curPassType
		if req.PassportType != "" {
			newType = req.PassportType
		}
		newNumber := curPassNumber
		if req.PassportNumber != "" {
			newNumber = req.PassportNumber
		}

		var exists bool
		// проверяем, что новый паспорт не принадлежит другому сотруднику
		if err = tx.QueryRowContext(ctx, scripts.IsPassportExistsExcludingID, newType, newNumber, passportID).Scan(&exists); err != nil {
			return err
		}
		if exists {
			return ErrPassportExist
		}

		// обновляем паспорт сотрудника
		if _, err = tx.ExecContext(ctx, scripts.UpdatePassportByID, passportID, newType, newNumber); err != nil {
			return err
		}
	}

	// обновляем самого сотрудника
	if _, err = tx.ExecContext(ctx, scripts.QueryUpdateEmployee, name, surname, phone, companyID, departmentID, req.ID); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && string(pqErr.Code) == "23505" {

			return ErrPhoneExist
		}
		return err
	}

	return tx.Commit()
}
