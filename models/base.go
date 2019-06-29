/*
 * Init function to connect to the DB
 */
package models

import (
	"fmt"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Initialising the DB")

	// Load credentials
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	// Open DB Connection
	db, err := sql.Open("mysql", "%s:%s@tcp(127.0.0.1:3306)/%s", username, password, dbName)
	if err != nil {
		panic(err.Error())
	}

	// Defer close till connection has finisched
	defer db.Close()
}
