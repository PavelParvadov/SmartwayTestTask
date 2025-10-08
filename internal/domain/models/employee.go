package models

type Employee struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Phone      string     `json:"phone"`
	CompanyId  int        `json:"company_id"`
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}
