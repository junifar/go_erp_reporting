package connection

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "193.168.195.28"
	port     = 5432
	user     = "integreat"
	password = "integreat"
	dbname   = "integreat"

	hostErp     = "192.168.0.100"
	portErp     = 5432
	userErp     = "openerp7"
	passwordErp = "openerp"
	dbnameErp   = "prasetia_dwidharma"
)

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ConnectErp() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostErp, portErp, userErp, passwordErp, dbnameErp)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
