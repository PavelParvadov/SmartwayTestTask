package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/pavelParvadov/SmartwayTask/internal/repository/postgres/scripts"
)

func (r *EmployeeRepositoryImpl) DeleteEmployee(ctx context.Context, employeeId int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var passportID sql.NullInt64
	// получаем passport_id сотрудника
	err = tx.QueryRowContext(ctx, scripts.QueryFindPassportID, employeeId).Scan(&passportID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrEmployeeNotFound
		}
		return err
	}

	// удаляем сотрудника
	if _, err = tx.ExecContext(ctx, scripts.QueryDeleteEmployee, employeeId); err != nil {
		return err
	}

	// если у сотрудника был паспорт, то его тоже удаляем
	if passportID.Valid {
		if _, err = tx.ExecContext(ctx, scripts.QueryDeletePassport, passportID.Int64); err != nil {
			return err
		}
	}

	return tx.Commit()
}
