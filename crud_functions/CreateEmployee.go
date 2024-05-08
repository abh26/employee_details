package crud_functions

import (
	"log"
)

// CreateEmployee creates a new employee record in the database
func (d *DatabaseManager) CreateEmployee(EmployeeName string, EmployeeId int, EmployeePosition string, EmployeeSalary float64) string {
	d.DbLock.Lock()
	defer d.DbLock.Unlock()

	create := Employee{
		Name:     EmployeeName,
		ID:       EmployeeId,
		Position: EmployeePosition,
		Salary:   EmployeeSalary,
	}
	result := d.Db.Table("Employee").Create(&create)
	if result.Error != nil {
		log.Println("Error found is " + result.Error.Error())
		return "Failed to store employee details"
	}
	return "Employee details saved successfully in the database"
}
