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
JOIN departments d ON d.id = e.department_id
LEFT JOIN passports p ON p.id = e.passport_id
WHERE e.department_id = $1 AND e.company_id = $2;


