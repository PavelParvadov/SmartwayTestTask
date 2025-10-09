CREATE TABLE IF NOT EXISTS companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS passports (
    id SERIAL PRIMARY KEY,
    type VARCHAR(15) NOT NULL,
    number VARCHAR(15) NOT NULL,
    CONSTRAINT uq_passports_type_number UNIQUE (type, number)
);

CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    department_id INT REFERENCES departments(id) ON DELETE SET NULL,
    passport_id INT UNIQUE REFERENCES passports(id) ON DELETE SET NULL
);

-- helpful indexes
CREATE INDEX IF NOT EXISTS idx_employees_company_id ON employees(company_id);
CREATE INDEX IF NOT EXISTS idx_employees_department_id ON employees(department_id);
CREATE INDEX IF NOT EXISTS idx_passports_type_number ON passports(type, number);
CREATE INDEX IF NOT EXISTS idx_employees_passport_id ON employees(passport_id);

-- seed data (optional)
INSERT INTO companies (name) VALUES
     ('Smartway'),
     ('T-bank');

INSERT INTO passports (type, number) VALUES
    ('5621', '323232'),
    ('5622', '611111'),
    ('5623', '777777'),
    ('5624', '332222');

INSERT INTO departments (name, phone, company_id) VALUES
     ('IT', '+7 953 022 15 68', 1),
     ('HR', '8-800-555-35-35', 2),
     ('Support', '89005553535', 1);

INSERT INTO employees (name, surname, phone, company_id, passport_id, department_id) VALUES
      ('Павел',     'Парвадов', '+7-202-555-0140', 1, 1, 1),
      ('Ксения',   'Доронина', '+7-202-555-0172', 2, 2, 2),
      ('Дональд',     'Трамп',    '+7-202-555-0114', 1, 3, 3),
      ('Игорь', 'Стрелков',    '+7-202-555-0183', 2, 4, 2);


