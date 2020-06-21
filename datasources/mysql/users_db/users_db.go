package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const (
	mysqlRootPassword = "MYSQL_ROOT_PASSWORD"
)

var (
	// Client -
	Client        *sql.DB
	mysqlPassword = os.Getenv(mysqlRootPassword)
)

// GetMySQLPassword -
func GetMySQLPassword() string {
	return mysqlPassword
}

func init() {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		GetMySQLPassword(),
		"127.0.0.1:3306",
		"users_db",
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
