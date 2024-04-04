package models

import (
	"context"
	"database/sql"
	"time"
)

type EmployeeModeler interface {
	Get(id string) (*Employee, error)
	GetAll() ([]*Employee, error)
	Insert(employee *Employee) error
}

type Employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type EmployeeModel struct {
	DB *sql.DB
}

func (e EmployeeModel) Get(id string) (*Employee, error) {
	query := `
		SELECT id, name, email
		FROM employee
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var employee Employee
	err := e.DB.QueryRowContext(ctx, query, id).Scan(&employee.ID, &employee.Name, &employee.Email)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (e EmployeeModel) GetAll() ([]*Employee, error) {
	query := `
		SELECT id, name, email
		FROM employee
		ORDER BY id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// time.Sleep(5 * time.Second)

	rows, err := e.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []*Employee{}
	for rows.Next() {
		var employee Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Email)
		if err != nil {
			return nil, err
		}
		employees = append(employees, &employee)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (e EmployeeModel) Insert(employee *Employee) error {
	query := `
		INSERT INTO employee (name, email)
		VALUES ($1, $2)
		RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{employee.Name, employee.Email}

	return e.DB.QueryRowContext(ctx, query, args...).Scan(&employee.ID)
}
