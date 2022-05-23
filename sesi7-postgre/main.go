package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID       int    `json:"int"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "tasrifin"
	DB_PASSWORD = "tasrifin"
	DB_NAME     = "dbtest"
)

func main() {
	db, err := connectDB()

	if err != nil {
		panic(err)
	}

	//CREATE EMPLOYEE
	emp := Employee{
		Email:    "hsh@kw.com",
		FullName: "test insert",
		Age:      21,
		Division: "Div",
	}

	err = createEmployee(db, &emp)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Database Connected!")

	//UPDATE EMPLOYEE
	empUpdate := Employee{
		Email:    "update2@kw.com",
		FullName: "test update2",
		Age:      23,
		Division: "Div",
	}

	err = updateEmployee(db, 2, &empUpdate)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Employee Updated")

	//DELETED EMPLOYEE
	err = deleteEmployee(db, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Employee Deleted")

	emps, err := getAllEmployees(db)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, emp := range *emps {
		emp.Print()
	}

	fmt.Println("====== Get Employee by id 2 ======")
	employee, err := getEmployeeById(db, 2)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}
	employee.Print()
}

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname =%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	// defer db.Close()
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func getAllEmployees(db *sql.DB) (*[]Employee, error) {
	query := `
		SELECT id, full_name, email, age, division from employees
	`

	statement, err := db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()
	var employees []Employee

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employee Employee
		err := rows.Scan(
			&employee.ID,
			&employee.FullName,
			&employee.Email,
			&employee.Age,
			&employee.Division,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil
}

func (e *Employee) Print() {
	fmt.Println("----------------------------------------------")
	fmt.Println("ID \t\t:", e.ID)
	fmt.Println("FullName \t:", e.FullName)
	fmt.Println("Email \t\t:", e.Email)
	fmt.Println("Age \t\t:", e.Age)
	fmt.Println("Division \t:", e.Division)
	fmt.Println("----------------------------------------------")
}

func createEmployee(db *sql.DB, request *Employee) error {
	query := `
		INSERT INTO employees(full_name, email, age, division)
		VALUES($1,$2,$3,$4)
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(request.FullName, request.Email, request.Age, request.Division)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()

}

func deleteEmployee(db *sql.DB, id uint) error {
	query := `
		DELETE from employees
		WHERE id = $1;
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func updateEmployee(db *sql.DB, id uint, request *Employee) error {
	query := `
		UPDATE employees
		SET full_name = $2, email = $3, division = $4, age = $5
		WHERE id = $1;
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id, request.FullName, request.Email, request.Division, request.Age)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func getEmployeeById(db *sql.DB, id int) (*Employee, error) {
	query := `
		SELECT id, full_name, email, age, division
		FROM employees
		WHERE id=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var emp Employee

	err = row.Scan(
		&emp.ID, &emp.FullName, &emp.Email, &emp.Age, &emp.Division,
	)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}
