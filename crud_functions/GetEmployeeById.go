package crud_functions

import (
	"fmt"
	"log"
)

// GetEmployeeById retrieves an employee record from the database by ID
func (d *DatabaseManager) GetEmployeeById(EmployeeId int) (string, error) {
	d.DbLock.Lock()
	defer d.DbLock.Unlock()

	var EmployeeDetails Employee
	result := d.Db.Table("Employee").Where("ID=?", EmployeeId).Find(&EmployeeDetails)
	if result.Error != nil {
		log.Println("Error found is " + result.Error.Error())
		return "Failed to retrieve employee details", result.Error
	}
	return fmt.Sprintf("Employee details for the given Id %d is : %s ", EmployeeId, fmt.Sprint(EmployeeDetails)), nil
}
