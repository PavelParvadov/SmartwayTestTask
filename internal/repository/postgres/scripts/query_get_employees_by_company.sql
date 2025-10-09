SELECT e.id,
       e.name,
       e.surname,
       e.phone,
       e.company_id,
       e.department_id,
       COALESCE(p.type, ''),
       COALESCE(p.number, ''),
       COALESCE(d.name, ''),
       COALESCE(d.phone, '')
FROM employees e
LEFT JOIN passports p ON p.id = e.passport_id
LEFT JOIN departments d ON d.id = e.department_id
WHERE e.company_id = $1;


