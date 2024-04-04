package models

import "database/sql"

type Models struct {
	Employee EmployeeModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Employee: EmployeeModel{DB: db},
	}
}
