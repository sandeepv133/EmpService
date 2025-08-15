package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sandeepv133/EmpService/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterEmployeeRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
