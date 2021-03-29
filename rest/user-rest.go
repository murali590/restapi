package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/murali590/restapi/controller"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// SetupRouter - sets the REST routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	v1 := r.Group("/v1")
	{
		v1.GET("user", controller.Getusers)
		v1.POST("user", controller.Createuser)
		v1.GET("user/:id", controller.Getuser)
		v1.PUT("user/:id", controller.Updateuser)
		v1.DELETE("user/:id", controller.Deleteuser)
	}

	return r
}
