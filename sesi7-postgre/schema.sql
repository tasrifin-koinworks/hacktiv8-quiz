--CREATE TABLE : EMPLOYEE
CREATE TABLE employees(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(50) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INT NOT NULL,
    division VARCHAR(50) NOT NULL
)