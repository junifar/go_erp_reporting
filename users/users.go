package users

import (
	"erp_reporting/connection"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ReturnAllUsers(){
	var users Users
	var arrUser []Users

	db := connection.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name from USERS ORDER BY ID ASC")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.Name); err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println(users.Id, users.Name)
			arrUser = append(arrUser, users)
		}
	}
}