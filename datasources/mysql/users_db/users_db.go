package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/angadthandi/bookstore_utils-go/logger"
	"github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {
	var err error
	log.Println("init database connection...")

	// user:pwd@tcp(host:port)/schema?charset=utf8
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		// failed to connect to DB...
		log.Fatalf("unable to connect to mysql db error: %v", err)
	}

	err = Client.Ping()
	if err != nil {
		log.Fatalf("unable to connect to mysql db error: %v", err)
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")

	return
}
