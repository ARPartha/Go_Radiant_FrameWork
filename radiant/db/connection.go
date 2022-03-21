package db

import (
	"fmt"
	"radiant/radiant/core"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func DBInit() {
	dbconf := core.GetDbConfig()
	connect_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",dbconf.DB_HOST,dbconf.DB_PORT,dbconf.DB_USERNAME, dbconf.DB_PASSWORD, dbconf.DB_NAME)
	db, err = gorm.Open("postgres", connect_string)
	//  defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}else{
		fmt.Println("Database successfully connected")
	}
	//  db.AutoMigrate(&models.User_info{})

}

func DbManager() *gorm.DB {
	return db
}
