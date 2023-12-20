// services/employee_service.go
package services

import (
	"hrms/internal/daos"
	"hrms/internal/models"
)

type EmployeeService struct {
    employeeDao *daos.EmployeeDao
}

func NewEmployeeService(employeeDao *daos.EmployeeDao) *EmployeeService {
    return &EmployeeService{employeeDao: employeeDao}
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) error {
    return s.employeeDao.Create(employee)
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
    return s.employeeDao.FindAll()
}

func (s *EmployeeService) GetEmployeeByID(id uint) (*models.Employee, error) {
    return s.employeeDao.FindByID(id)
}

func (s *EmployeeService) UpdateEmployee(employee *models.Employee) error {
    return s.employeeDao.Update(employee)
}

func (s *EmployeeService) DeleteEmployee(id uint) error {
    return s.employeeDao.Delete(id)
}
