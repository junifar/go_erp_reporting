package main

import (
	"erp_reporting/budget_realization"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("It's Worked")
	//users.ReturnAllUsers()
	//budget_realization.GetBudgetRealization()
	runServer()
}

func runServer() {
	router := mux.NewRouter()
	router.HandleFunc("/budget_realization/{tahun}/{dept_id}", budget_realization.GetBudgetRealization).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
