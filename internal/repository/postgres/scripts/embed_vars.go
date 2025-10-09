package scripts

import _ "embed"

var (
	//go:embed is_company_exists.sql
	IsCompanyExists string
	//go:embed is_passport_exists.sql
	IsPassportExists string
	//go:embed insert_passport.sql
	InsertPassport string
	//go:embed is_department_exists.sql
	IsDepartmentExists string
	//go:embed insert_employee.sql
	InsertEmployee string
	//go:embed query_find_passport_id.sql
	QueryFindPassportID string
	//go:embed query_delete_employee.sql
	QueryDeleteEmployee string
	//go:embed query_delete_passport.sql
	QueryDeletePassport string
	//go:embed update_passport_by_id.sql
	UpdatePassportByID string
	//go:embed is_passport_exists_excluding_id.sql
	IsPassportExistsExcludingID string
	//go:embed query_find_employee.sql
	QueryFindEmployee string
	//go:embed query_update_employee.sql
	QueryUpdateEmployee string
	//go:embed query_find_department_by_company.sql
	QueryFindDepartmentByCompany string
	//go:embed query_get_employees_by_department.sql
	QueryGetEmployeesByDepartment string
	//go:embed query_get_employees_by_company.sql
	QueryGetEmployeesByCompany string
)
