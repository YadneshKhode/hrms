// main.go
package main

import (
	"hrms/internal/controllers"
	"hrms/internal/daos"
	"hrms/internal/models"
	"hrms/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    // Connect to the database
    dsn := "roo:Password@tcp(127.0.0.1:3306)/hrms?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database")
    }

    // Migrate the schema
    db.AutoMigrate(&models.Employee{})

    // Create instances of DAO, Service, and Controller
    employeeDao := daos.NewEmployeeDao(db)
    employeeService := services.NewEmployeeService(employeeDao)
    employeeController := controllers.NewEmployeeController(employeeService)

    // Set up Fiber
    app := fiber.New()
	
	// Group routes with a common prefix
	employeesGroup := app.Group("/employees")

    // Define routes
	{
		employeesGroup.Post("", employeeController.CreateEmployee)
		employeesGroup.Get("", employeeController.GetAllEmployees)
		employeesGroup.Get("/:id", employeeController.GetEmployeeByID)
		employeesGroup.Put("/:id", employeeController.UpdateEmployee)
		employeesGroup.Delete("/:id", employeeController.DeleteEmployee)
	}
    // Start the server
    app.Listen(":3000")
}
