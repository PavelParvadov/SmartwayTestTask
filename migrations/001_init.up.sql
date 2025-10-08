
CREATE TABLE IF NOT EXISTS companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
    );


CREATE TABLE IF NOT EXISTS departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone VARCHAR(20),
    company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE
    );


CREATE TABLE IF NOT EXISTS passports (
    id SERIAL PRIMARY KEY,
    type VARCHAR(15) NOT NULL,
    number VARCHAR(15) NOT NULL
    );


CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    company_id INT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    department_id INT REFERENCES departments(id) ON DELETE SET NULL,
    passport_id INT UNIQUE REFERENCES passports(id) ON DELETE CASCADE
    );
