package model

type Employee struct {
	EmpId int `json:"empid"`
    Name string `json:"name"`
	Email string `json:"email"`
    Address Address `json:"address"`
}

const(
	EmpId = 0
	Name = 2
	Email = 6
)

func NewEmployee( empId int, name, email string, address *Address) *Employee {
    employee := new(Employee)
    employee.EmpId = empId
	employee.Name = name
	employee.Email = email
	employee.Address = *address
    return employee
}