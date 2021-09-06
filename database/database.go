package database

import (
	"database/sql"
	"fmt"

	"gojek.com/config"
	"gojek.com/logger"

	_ "github.com/go-sql-driver/mysql"
)

var primaryDBH *sql.DB

func GetConnection() *sql.DB {
	if primaryDBH != nil {
		return primaryDBH
	}

	host := config.Configs.DBHost
	user := config.Configs.DBUser
	password := config.Configs.DBPassword
	port := config.Configs.DBPort
	schema := config.Configs.DBSchema
	driver := config.Configs.DBDriver

	connectionString := fmt.Sprintf("%[1]s:%[2]s@tcp(%[3]s:%[4]s)/%[5]s", user, password, host, port, schema)
	connectionString += "?parseTime=true&charset=utf8"

	db, err := sql.Open(driver, connectionString)
	if err != nil {
		logger.Error("Failed connecting to database." + err.Error())
	}

	primaryDBH = db
	return primaryDBH
				
}