package main

import (
	"fmt"
)

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Full Time Employee: ID: %d, Name: %s, Salary: %d Tenge", f.ID, f.Name, f.Salary)
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Part Time Employee: ID: %d, Name: %s, Hourly Rate: %d Tenge, Hours Worked: %.2f", p.ID, p.Name, p.HourlyRate, p.HoursWorked)
}

type Company struct {
	employees map[string]Employee
}

func NewCompany() *Company {
	return &Company{
		employees: make(map[string]Employee),
	}
}

func (c *Company) AddEmployee(emp Employee) {
	switch e := emp.(type) {
	case FullTimeEmployee:
		c.employees[fmt.Sprintf("%d", e.ID)] = e
	case PartTimeEmployee:
		c.employees[fmt.Sprintf("%d", e.ID)] = e
	}
}

func (c *Company) ListEmployees() {
	if len(c.employees) == 0 {
		fmt.Println("No employees found.")
		return
	}
	for _, emp := range c.employees {
		fmt.Println(emp.GetDetails())
	}
}

func (c *Company) RemoveEmployee(id uint64) {
	key := fmt.Sprintf("%d", id)
	if _, exists := c.employees[key]; exists {
		delete(c.employees, key)
		fmt.Printf("Employee with ID %d removed.\n", id)
	} else {
		fmt.Printf("Employee with ID %d not found.\n", id)
	}
}

func ParseFullTimeEmployee() FullTimeEmployee {
	var id uint64
	var name string
	var salary uint32

	fmt.Print("Enter FullTime Employee ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter FullTime Employee Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter FullTime Employee Salary: ")
	fmt.Scan(&salary)

	return FullTimeEmployee{ID: id, Name: name, Salary: salary}
}

func ParsePartTimeEmployee() PartTimeEmployee {
	var id uint64
	var name string
	var hourlyRate uint64
	var hoursWorked float32

	fmt.Print("Enter PartTime Employee ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter PartTime Employee Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter PartTime Employee Hourly Rate: ")
	fmt.Scan(&hourlyRate)
	fmt.Print("Enter PartTime Employee Hours Worked: ")
	fmt.Scan(&hoursWorked)

	return PartTimeEmployee{ID: id, Name: name, HourlyRate: hourlyRate, HoursWorked: hoursWorked}
}

func main() {
	c := NewCompany()

	for {
		fmt.Println("\nEmployee Management System")
		fmt.Println("1. Add FullTime Employee")
		fmt.Println("2. Add PartTime Employee")
		fmt.Println("3. List Employees")
		fmt.Println("4. Remove Employee")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			emp := ParseFullTimeEmployee()
			c.AddEmployee(emp)
		case 2:
			emp := ParsePartTimeEmployee()
			c.AddEmployee(emp)
		case 3:
			c.ListEmployees()
		case 4:
			fmt.Print("Enter Employee ID to remove: ")
			var id uint64
			fmt.Scan(&id)
			c.RemoveEmployee(id)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
