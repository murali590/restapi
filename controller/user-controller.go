package controller

import (
	"log"
	"net/http"

	"github.com/Eswaraa/inflix/service"
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
	var users model.user
	// var mapAlbum map[string]interface{}
	if c.BindJSON(&user) == nil {
		log.Println(album.Title)
		log.Println(album.Release)
		log.Println(album.Production)
		log.Println(album.Director)
		log.Println(album.IsPremium)
	} else {
		log.Println("Error in binding")
	}

	err := service.CreateAlbum(&album)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, album)
	}
}
