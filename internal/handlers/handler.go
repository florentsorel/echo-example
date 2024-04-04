package handlers

import (
	"echo-example/api/internal/services"
)

type handler struct {
	EmployeeHandler EmployeeHandler
}

func new(s *services.Service) *handler {
	return &handler{
		EmployeeHandler: EmployeeHandler{
			EmployeeService: &s.Employee,
		},
	}
}
