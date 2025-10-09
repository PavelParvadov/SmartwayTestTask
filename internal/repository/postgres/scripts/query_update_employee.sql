UPDATE employees
SET name = $1,
    surname = $2,
    phone = $3,
    company_id = $4,
    department_id = $5
WHERE id = $6;



