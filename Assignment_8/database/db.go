package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
const DB_USERNAME = "EdngiFqCZd"
const DB_PASSWORD = "if44fVauir"
const DB_NAME = "EdngiFqCZd"
const DB_HOST = "remotemysql.com"
const DB_PORT = "3306"

var Db *gorm.DB
func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() (*gorm.DB) {
	var err error
	dbHostURL := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dbHostURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
