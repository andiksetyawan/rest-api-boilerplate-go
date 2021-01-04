package main

import (
	"flag"
	"github.com/andiksetyawan/rest-api-boilerplate-go/controller"
	_ "github.com/andiksetyawan/rest-api-boilerplate-go/db"
	"github.com/andiksetyawan/rest-api-boilerplate-go/middleware"
	_ "github.com/andiksetyawan/rest-api-boilerplate-go/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	bindAddr := flag.String("bind", ":8080", "bind addr")
	flag.Parse()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	r.GET("/file/:bucket/:id", controller.DownloadFile)

	r.POST("/login", controller.Login)
	r.POST("/signup", controller.SignUp)

	//r.Use(midlewares.IsAuth())
	r.GET("/user", middleware.IsAuth(), controller.GetUser)
	r.Run(*bindAddr)
}
