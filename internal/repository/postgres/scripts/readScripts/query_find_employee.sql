SELECT e.name,
       e.surname,
       e.phone,
       e.company_id,
       p.type,
       p.number,
       d.name,
       d.phone,
       e.department_id
FROM employees e
LEFT JOIN passports p ON p.id = e.passport_id
LEFT JOIN departments d ON d.id = e.department_id
WHERE e.id = $1;




