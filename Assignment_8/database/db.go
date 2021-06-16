package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/magiconair/properties"
	"os"
	"path/filepath"
	"log"
)


var Db *gorm.DB


func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}



func connectDB() (*gorm.DB) {
	var prop = readDBPropertyFile()	
	var db_username = prop.MustGet("db.username")
	var db_password = prop.MustGet("db.password")
	var db_name = prop.MustGet("db.name")
	var db_host = prop.MustGet("db.host")
	var db_port = prop.MustGet("db.port")
	var err error
	dbHostURL := db_username +":"+ db_password +"@tcp"+ "(" + db_host + ":" + db_port +")/" + db_name + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dbHostURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : error=", err)
		return nil
	}

	return db
}


func readDBPropertyFile() *properties.Properties {
	path, _ := os.Getwd()
	var PropertyFile= []string{filepath.Join(path, "db.properties")}
	var prop, err = properties.LoadFiles(PropertyFile, properties.UTF8, true)
	if err != nil {
		log.Fatal("Error in reading property file error=", err)
		return nil
	}
	return prop
 }

