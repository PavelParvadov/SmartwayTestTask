package postgres

import (
	"context"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
	"github.com/pavelParvadov/SmartwayTask/internal/repository/postgres/scripts"
)

func (r *EmployeeRepositoryImpl) GetEmployeesByCompanyID(ctx context.Context, companyId int) ([]models.Employee, error) {
	var exists bool
	// проверяем существует ли компания
	if err := r.db.QueryRowContext(ctx, scripts.IsCompanyExists, companyId).Scan(&exists); err != nil {
		return nil, err
	}
	if !exists {
		return []models.Employee{}, nil
	}

	// получаем сотрудников компании
	rows, err := r.db.QueryContext(ctx, scripts.QueryGetEmployeesByCompany, companyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := make([]models.Employee, 0)
	for rows.Next() {
		var e models.Employee
		if err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.Surname,
			&e.Phone,
			&e.CompanyId,
			&e.DepartmentId,
			&e.Passport.Type,
			&e.Passport.Number,
			&e.Department.Name,
			&e.Department.Phone,
		); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}
