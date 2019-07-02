/*
 * Init function to connect to the DB
 */

package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {

	//TODO add env variables
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
	dbName := "golang-db"
	dbHost := "db"

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName)
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
