// controllers/employee_controller.go
package controllers

import (
	"hrms/internal/models"
	"hrms/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
    employeeService *services.EmployeeService
}

func NewEmployeeController(employeeService *services.EmployeeService) *EmployeeController {
    return &EmployeeController{employeeService: employeeService}
}

func (c *EmployeeController) CreateEmployee(ctx *fiber.Ctx) error {
    var employee models.Employee
    if err := ctx.BodyParser(&employee); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    if err := c.employeeService.CreateEmployee(&employee); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.JSON(employee)
}

func (c *EmployeeController) GetAllEmployees(ctx *fiber.Ctx) error {
    employees, err := c.employeeService.GetAllEmployees()
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return ctx.JSON(employees)
}

func (c *EmployeeController) GetEmployeeByID(ctx *fiber.Ctx) error {
    id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee ID"})
    }

    employeeID := uint(id)
    employee, err := c.employeeService.GetEmployeeByID(employeeID)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.JSON(employee)
}

func (c *EmployeeController) UpdateEmployee(ctx *fiber.Ctx) error {
    id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee ID"})
    }

    employeeID := uint(id)
    var updatedEmployee models.Employee
    if err := ctx.BodyParser(&updatedEmployee); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    existingEmployee, err := c.employeeService.GetEmployeeByID(employeeID)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    existingEmployee.Name = updatedEmployee.Name
    existingEmployee.Email = updatedEmployee.Email
    existingEmployee.Address = updatedEmployee.Address

    if err := c.employeeService.UpdateEmployee(existingEmployee); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.JSON(existingEmployee)
}

func (c *EmployeeController) DeleteEmployee(ctx *fiber.Ctx) error {
    id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee ID"})
    }

    employeeID := uint(id)
    if err := c.employeeService.DeleteEmployee(employeeID); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.SendStatus(fiber.StatusNoContent)
}
