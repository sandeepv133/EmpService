package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sandeepv133/EmpService/pkg/config"
)

var db *gorm.DB

type Employee struct {
	gorm.Model
	First_name  string `json:"first_name"`
	Middle_name string `json:"middle_name"`
	Last_name   string `json:"last_name"`
	// Dob                  time.Time `json:"dob"`
	// Hire_date            time.Time `json:"hire_date"`
	Stat                 string `json:"stat"`
	Email                string `json:"email"`
	Phone_number         string `json:"phone_number"`
	Profile_pic_location string `json:"profile_pic_location"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Employee{})
}

func (e *Employee) CreateEmployee() *Employee {
	db.NewRecord(e)
	db.Create(&e)
	return e
}

func GetAllEmployees() []Employee {
	var Employees []Employee
	db.Find(&Employees)
	return Employees
}

func GetEmployeeById(Id int64) (*Employee, *gorm.DB) {
	var getEmployee Employee
	db := db.Where("ID=?", Id).Find(&getEmployee)
	return &getEmployee, db

}

func DeleteEmployee(Id int64) Employee {
	var employee Employee
	db.Where("id = ?", Id).First(&employee)
	if employee.Stat == "Active" {
		db.Model(&employee).Update("stat", "Inactive") // <-- writes to DB
		employee.Stat = "Inactive"                     // keep in-memory copy in sync
	}
	db.Where("Id=?", Id).Delete(employee)
	return employee

}

func UpdateEmployee(Id int64) Employee {
	var employee Employee
	db.Where("Id=?", Id).Update(employee)
	return employee
}
