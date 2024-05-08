package crud_functions

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (d *DatabaseManager) TestCreateEmployee_Success(t *testing.T) {
	// Mock DB using SQLite in-memory database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	// Create the Employee table
	db.AutoMigrate(&Employee{})

	// Call the function with valid input
	result := d.CreateEmployee("Abhishek Rai", 1, "Manager", 50000)
	assert.Equal(t, "Employee details saved successfully in the database", result)
}

func (d *DatabaseManager) TestCreateEmployee_Failure(t *testing.T) {
	// Mock DB using SQLite in-memory database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	// Create the Employee table
	db.AutoMigrate(&Employee{})

	// Call the function with invalid input (e.g., duplicate EmployeeId)
	d.CreateEmployee("Aashil", 1, "Manager", 50000)
	result := d.CreateEmployee( "Girish", 1, "Developer", 60000)
	assert.Contains(t, result, "Failed to store employee details")
}
func (d *DatabaseManager) TestDeleteEmployee_Success(t *testing.T) {
	// Mock DB using SQLite in-memory database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	// Create the Employee table
	db.AutoMigrate(&Employee{})

	// Add an employee to the database
	db.Create(&Employee{Name: "aashil", ID: 1, Position: "Manager", Salary: 50000})

	// Call the function to delete the employee
	result := d.DeleteEmployee( 1)
	assert.Equal(t, "Employee details deleted successfully", result)
}

func (d *DatabaseManager) TestDeleteEmployee_Failure(t *testing.T) {
	// Mock DB using SQLite in-memory database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	// Create the Employee table
	db.AutoMigrate(&Employee{})

	// Call the function with an employee ID that does not exist
	result := d.DeleteEmployee(1)
	assert.Contains(t, result, "Failed to delete employee details")
}

func (d *DatabaseManager) TestGetEmployeeById_Success(t *testing.T) {
	// Mock DB using SQLite in-memory database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	// Create the Employee table
	db.AutoMigrate(&Employee{})

	// Add an employee to the database
	db.Create(&Employee{Name: "ashil", ID: 1, Position: "Manager", Salary: 50000})

	// Call the function to retrieve the employee by ID
	result,fetcherr:= d.GetEmployeeById(1)
	if fetcherr != nil {
		log.Println("error : ", err)
	}
	assert.Contains(t, result, "Employee details fetched successfully")
}

func (d *DatabaseManager) TestGetEmployeeById_Failure(t *testing.T) {
	// Mock DB using SQLite in-memory database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	// Create the Employee table
	db.AutoMigrate(&Employee{})

	// Call the function with an employee ID that does not exist
	result,fetcherr := d.GetEmployeeById(1)
	if fetcherr != nil {
		log.Println("error : ", err)
	}
	assert.Contains(t, result, "Failed to retrieve employee details")
}

func (d *DatabaseManager) TestUpdateEmployee_Success(t *testing.T) {
	// Mock DB using SQLite in-memory database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	// Create the Employee table
	db.AutoMigrate(&Employee{})

	// Add an employee to the database
	db.Create(&Employee{Name: "aashil", ID: 1, Position: "Manager", Salary: 50000})

	// Call the function to update the employee details
	result:= d.UpdateEmployee("Smith", 1, "Developer", 60000)
	if err != nil {
		log.Println("error : ", err)
	}
	assert.Contains(t, result, "Employee details updated successfully")
}

func (d *DatabaseManager) TestUpdateEmployee_Failure(t *testing.T) {
	// Mock DB using SQLite in-memory database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	// Create the Employee table
	db.AutoMigrate(&Employee{})

	// Call the function with an employee ID that does not exist
	result := d.UpdateEmployee("Smith", 1, "Developer", 60000)
	assert.Contains(t, result, "Failed to update employee details")
}
