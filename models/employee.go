package models

type Employee struct {
	EmployeeId string `json:"emplyee_id,omitempty" bson:"employee_id"`
	Name       string `json:"name,omitempty" bson:"name"`
	Department string `json:"department,omitempty" bson:"department"`
}
