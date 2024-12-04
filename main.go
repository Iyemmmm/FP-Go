package main

import (
	// "fmt"
	"net/http"
	"resto/controller"
	"resto/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.ConnectDatabase()
}

func main(){
	r:=gin.Default()

	r.LoadHTMLGlob("view/*")

	pathImg:="./upload"
	r.Static("/static", "./static")
	r.Static("/upload", "./upload")

	
	r.GET("/",controller.Welcome)
	r.GET("/about",func(c *gin.Context){
		c.HTML(http.StatusOK,"about.html",nil)
	})
	
	r.GET("/menu",controller.ViewMenu)
	r.POST("/menu",controller.AddCart)

	r.GET("/reservation",controller.ViewReservation)

	r.GET("/register",controller.ViewRegister)
	r.POST("/register",controller.CreateUser)

	r.GET("/login",controller.ViewLogin)
	r.POST("/login",controller.LoginUser)

	r.GET("/addmenu",func(c *gin.Context) {
		c.HTML(http.StatusOK,"addmenu.html",nil)
	})
	r.POST("/addmenu",func(c *gin.Context) {
		controller.AddMenu(c,pathImg)
	})

	r.GET("/cart",controller.ViewCart)

	r.Run(":8080")
}