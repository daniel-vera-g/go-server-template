/*
 * Init function to connect to the DB
 */
package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {

	//TODO add enc variables
	// e := godotenv.Load()
	// if e != nil {
	// fmt.Print(e)
	// }

	// username := os.Getenv("db_user")
	// password := os.Getenv("db_pass")
	// dbName := os.Getenv("db_name")
	// dbHost := os.Getenv("db_host")

	username := "root"
	password := "password"
	dbName := "notes"
	dbHost := "172.18.0.2"

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Note{})
}

func GetDB() *gorm.DB {
	return db
}
