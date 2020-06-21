package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const (
	mysqlUsersUsername = "MYSQL_ROOT_USERNAME"
	mysqlRootPassword  = "MYSQL_ROOT_PASSWORD"
	mysqlUsersHost     = "MYSQL_ROOT_HOST"
	mysqlUsersSchema   = "MYSQL_ROOT_SCHEMA"
)

var (
	// Client -
	Client        *sql.DB
	mysqlUsername = os.Getenv(mysqlUsersUsername)
	mysqlPassword = os.Getenv(mysqlRootPassword)
	mysqlHost     = os.Getenv(mysqlUsersHost)
	mysqlSchema   = os.Getenv(mysqlUsersSchema)
)

// GetMySQLPassword - getting mysql password
func GetMySQLPassword() string {
	return mysqlPassword
}

// GetMySQLUsername -
func GetMySQLUsername() string {
	return mysqlUsername
}

// GetMySQLHost -
func GetMySQLHost() string {
	return mysqlHost
}

// GetMySQLSchema -
func GetMySQLSchema() string {
	return mysqlSchema
}

func init() {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		GetMySQLUsername(),
		GetMySQLPassword(),
		GetMySQLHost(),
		GetMySQLSchema(),
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
