// daos/employee_dao.go
package daos

import (
	"hrms/internal/models"

	"gorm.io/gorm"
)

type EmployeeDao struct {
    db *gorm.DB
}

func NewEmployeeDao(db *gorm.DB) *EmployeeDao {
    return &EmployeeDao{db: db}
}

func (dao *EmployeeDao) Create(employee *models.Employee) error {
    return dao.db.Create(employee).Error
}

func (dao *EmployeeDao) FindAll() ([]models.Employee, error) {
    var employees []models.Employee
    if err := dao.db.Find(&employees).Error; err != nil {
        return nil, err
    }
    return employees, nil
}

func (dao *EmployeeDao) FindByID(id uint) (*models.Employee, error) {
    var employee models.Employee
    if err := dao.db.First(&employee, id).Error; err != nil {
        return nil, err
    }
    return &employee, nil
}

func (dao *EmployeeDao) Update(employee *models.Employee) error {
    return dao.db.Save(employee).Error
}

func (dao *EmployeeDao) Delete(id uint) error {
    return dao.db.Delete(&models.Employee{}, id).Error
}
