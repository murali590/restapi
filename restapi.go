package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/murali590/restapi/config"
	"github.com/murali590/restapi/rest"
)

var err error

func main() {

	// Creating a connection to the database
	config.DB, err = gorm.Open("mysql", config.DbConnectionString(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Error Connecting to mydb DB: ", err)
	}

	defer config.DB.Close()

	// run the migrations: album struct
	config.DB.AutoMigrate(&model.user{})

	//setup routes
	r := rest.SetupRouter()

	// running
	r.Run()

}
