package crud_functions

import (
	"log"
)

// UpdateEmployee updates an employee record in the database
func (d *DatabaseManager) UpdateEmployee(EmployeeName string, UpdateId int, EmployeePosition string, EmployeeSalary float64) string {
	d.DbLock.Lock()
	defer d.DbLock.Unlock()

	updateDetails := Employee{
		Name:     EmployeeName,
		ID:       UpdateId,
		Position: EmployeePosition,
		Salary:   EmployeeSalary,
	}
	result := d.Db.Table("Employee").Where("ID=?", UpdateId).Updates(&updateDetails)
	if result.Error != nil {
		log.Println("Error found is " + result.Error.Error())
		return "Failed to update employee details"
	}
	return "Employee details updated successfully in the database"
}
