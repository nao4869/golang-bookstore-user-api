package usersdb

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	// Client -
	Client *sql.DB
	mysqlRootPassword = "MYSQL_ROOT_PASSWORD"
)

func init() {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"mysqlRootPassword",
		"127.0.0.1:3306",
		"usersdb",
	)

	var err error
	Client, err := sql.Open(
		"mysql",
		dataSourceName,
	)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully comfigured")
}