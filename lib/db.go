package lib

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var config = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "root",
	Password: "root",
	Database: "membros",
}

// Sess is variable to make database connection
var Sess db.Database

func init() {
	var err error
	Sess, err = mysql.Open(config)

	if err != nil {
		log.Fatal(err.Error())
	}
}
