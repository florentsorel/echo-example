package services

import "echo-example/api/internal/models"

type EmployeeServicer interface {
	Find(id string) (*models.Employee, error)
	FindAll() ([]*models.Employee, error)
	Create(employee *models.Employee) error
}

type EmployeeService struct {
	model models.EmployeeModel
}

func (es *EmployeeService) Find(id string) (*models.Employee, error) {
	return es.model.Get(id)
}

func (es *EmployeeService) FindAll() ([]*models.Employee, error) {
	return es.model.GetAll()
}

func (es *EmployeeService) Create(employee *models.Employee) error {
	return es.model.Insert(employee)
}
