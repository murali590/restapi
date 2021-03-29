package controller

import (
	"log"
	"net/http"

	"github.com/Eswaraa/restapi/service"
	"github.com/gin-gonic/gin"
)

//Getusera - List all the Albums
func Getusers(c *gin.Context) {
	var users []model.user
	err := service.GetAllusers(&users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
		log.Printf("Users List: %v", users)
	}
}

//Createuser - Create an user
func Createuser(c *gin.Context) {
	var us model.user
	// var mapuser map[string]interface{}
	if c.BindJSON(&us) == nil {
		log.Println(us.userid)
		log.Println(us.password)
		log.Println(us.name)
		log.Println(us.date_of_birth)
		log.Println(us.address)
	} else {
		log.Println("Error in binding")
	}

	err := service.Createuser(&us)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, us)
	}
}
