package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sandeepv133/EmpService/pkg/models"
	"github.com/sandeepv133/EmpService/pkg/utils"
)

var NewEmployee models.Employee

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	NewEmployees := models.GetAllEmployees()
	res, _ := json.Marshal(NewEmployees)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeId := vars["employeeId"]
	Id, err := strconv.ParseInt(employeeId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	employeeDetails, _ := models.GetEmployeeById(Id)
	res, _ := json.Marshal(employeeDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	CreateEmployee := &models.Employee{}
	utils.ParseBody(r, CreateEmployee)
	e := CreateEmployee.CreateEmployee()
	res, _ := json.Marshal(e)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeId := vars["employeeId"]
	Id, err := strconv.ParseInt(employeeId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	employee := models.DeleteEmployee(Id)
	res, _ := json.Marshal(employee)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var updateEmployee = &models.Employee{}
	utils.ParseBody(r, updateEmployee)
	vars := mux.Vars(r)
	employeeId := vars["employeeId"]
	Id, err := strconv.ParseInt(employeeId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	employeeDetails, db := models.GetEmployeeById(Id)
	if updateEmployee.First_name != "" {
		employeeDetails.First_name = updateEmployee.First_name
	}
	if updateEmployee.Last_name != "" {
		employeeDetails.Last_name = updateEmployee.Last_name
	}
	if updateEmployee.Middle_name != "" {
		employeeDetails.Middle_name = updateEmployee.Middle_name
	}
	if updateEmployee.Dob.IsZero() {
		employeeDetails.Dob = updateEmployee.Dob
	}
	if updateEmployee.Hire_date.IsZero() {
		employeeDetails.Hire_date = updateEmployee.Hire_date
	}
	if updateEmployee.Created_date_time.IsZero() {
		employeeDetails.Created_date_time = updateEmployee.Created_date_time
	}
	if updateEmployee.Stat != "" {
		employeeDetails.Stat = updateEmployee.Stat
	}
	if updateEmployee.Email != "" {
		employeeDetails.Email = updateEmployee.Email
	}
	if updateEmployee.Phone_number != "" {
		employeeDetails.Phone_number = updateEmployee.Phone_number
	}
	if updateEmployee.Profile_pic_location != "" {
		employeeDetails.Profile_pic_location = updateEmployee.Profile_pic_location
	}
	db.Save(&employeeDetails)
	res, _ := json.Marshal(employeeDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
