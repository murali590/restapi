package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murali590/restapi/service"
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

//Getuser - Get the user detail for a given user id
func Getuser(c *gin.Context) {
	id := c.Params.ByName("id")
	var us model.user
	err := service.Getuser(&us, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, us)
	}
}

// Updateuser - Update an existing user
func Updateuser(c *gin.Context) {
	var us model.user
	id := c.Params.ByName("id")
	err := service.Getuser(&us, id)
	if err != nil {
		c.JSON(http.StatusNotFound, us)
	}
	c.BindJSON(&us)
	err = service.Updateuser(&us, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, us)
	}
}

//Deleteuser - Delete an user
func Deleteuser(c *gin.Context) {
	var us model.user
	id := c.Params.ByName("id")
	err := service.Deleteuser(&us, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
