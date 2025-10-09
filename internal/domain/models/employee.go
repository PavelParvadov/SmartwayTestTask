package models

type Employee struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Surname      string     `json:"surname"`
	Phone        string     `json:"phone"`
	CompanyId    int        `json:"company_id"`
	DepartmentId int        `json:"department_id"`
	Passport     Passport   `json:"passport"`
	Department   Department `json:"department"`
}
