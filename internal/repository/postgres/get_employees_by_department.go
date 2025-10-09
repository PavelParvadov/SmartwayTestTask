package postgres

import (
	"context"

	"github.com/pavelParvadov/SmartwayTask/internal/domain/models"
	"github.com/pavelParvadov/SmartwayTask/internal/repository/postgres/scripts"
)

func (r *EmployeeRepositoryImpl) GetEmployeeByDepartmentID(ctx context.Context, DepartmentId int, CompanyId int) ([]models.Employee, error) {
	var exists bool
	// проверяем принадлежит ли отдел компании
	if err := r.db.QueryRowContext(ctx, scripts.QueryFindDepartmentByCompany, DepartmentId, CompanyId).Scan(&exists); err != nil {
		return nil, err
	}
	if !exists {
		return []models.Employee{}, nil
	}

	// получаем сотрудников отдела внутри компании
	rows, err := r.db.QueryContext(ctx, scripts.QueryGetEmployeesByDepartment, DepartmentId, CompanyId)
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
