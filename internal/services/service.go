package services

import "echo-example/api/internal/models"

type Service struct {
	Employee EmployeeService
}

func New(m models.Models) *Service {
	return &Service{
		Employee: EmployeeService{
			model: m.Employee,
		},
	}
}
