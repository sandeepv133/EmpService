package routes

import (
	"github.com/gorilla/mux"
	"github.com/sandeepv133/EmpService/pkg/controllers"
)

var RegisterEmployeeRoutes = func(router *mux.Router) {
	router.HandleFunc("/employee/", controllers.CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}/", controllers.GetEmployeeById).Methods("GET")
	router.HandleFunc("/employee/", controllers.GetAllEmployees).Methods("GET")
	router.HandleFunc("/employee/{id}/", controllers.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{id}/", controllers.DeleteEmployee).Methods("DELETE")
}
