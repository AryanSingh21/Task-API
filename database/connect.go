package connect

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var DB *gorm.DB
 
func ConnectDatabse() {
	db, err := gorm.Open(sqlite.Open("task_db.db"), &gorm.Config{})

	if err != nil {
		panic("Error while connecting to database")
	}
	err = db.AutoMigrate(&Task{})
	if err != nil {
		panic("Error while migrating to the database")
	}
	
	fmt.Println("Database connected successfully")

	DB = db

}