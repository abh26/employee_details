package crud_functions

import (
	"log"
)

// DeleteEmployee deletes an employee record from the database by ID
func (d *DatabaseManager) DeleteEmployee(EmployeeId int) string {
	d.DbLock.Lock()
	defer d.DbLock.Unlock()

	matchCondition := make(map[string]interface{})
	matchCondition["ID"] = EmployeeId
	result := d.Db.Table("Employee").Delete(&matchCondition)
	if result.Error != nil {
		log.Println("Error found is " + result.Error.Error())
		return "Failed to delete employee details"
	}
	return "Employee details deleted successfully"
}
