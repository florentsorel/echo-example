package handlers

import (
	"net/http"

	"echo-example/api/internal/models"
	"echo-example/api/internal/services"

	"github.com/labstack/echo/v4"
)

type EmployeeHandler struct {
	EmployeeService services.EmployeeServicer
}

func New(s *services.Service) *handler {
	return &handler{
		EmployeeHandler: EmployeeHandler{
			EmployeeService: &s.Employee,
		},
	}
}

func (h *EmployeeHandler) HandlerList(c echo.Context) error {
	employees, err := h.EmployeeService.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) HandlerGetByID(c echo.Context) error {
	id := c.Param("id")
	employee, err := h.EmployeeService.Find(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) HandlerCreate(c echo.Context) error {
	var employee models.Employee
	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(employee); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.EmployeeService.Create(&employee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, employee)
}
