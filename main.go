package main

import (
	// "fmt"
	"net/http"
	"resto/controller"
	"resto/initializers"
	"resto/middleware"

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

	
	r.GET("/",middleware.IsLogin,controller.Welcome)
	r.GET("/about",func(c *gin.Context){
		c.HTML(http.StatusOK,"about.html",nil)
	})
	
	r.GET("/menu",controller.ViewMenu)
	r.POST("/menu",middleware.RequireAuth,controller.AddCart)

	r.GET("/reservation",middleware.RequireAuth,controller.ViewReservation)
	r.POST("/reservation",middleware.RequireAuth,controller.AddReservation)

	r.GET("/register",controller.ViewRegister)
	r.POST("/register",controller.CreateUser)

	r.GET("/login",controller.ViewLogin)
	r.POST("/login",controller.LoginUser)

	// r.GET("/addmenu",middleware.AdminMiddleware(),func(c *gin.Context) {
	// 	c.HTML(http.StatusOK,"addmenu.html",nil)
	// })
	// r.POST("/addmenu",middleware.AdminMiddleware(),func(c *gin.Context) {
	// 	controller.AddMenu(c,pathImg)
	// })

	r.GET("/addmenu",middleware.AdminMiddleware,controller.ViewAddMenu)
	r.POST("/addmenu", middleware.AdminMiddleware, func(c *gin.Context) {
		controller.AddMenu(c, pathImg)
	})

	r.GET("/reservation-status",middleware.RequireAuth,controller.ViewStatus)
	r.GET("/cart",middleware.RequireAuth,controller.ViewCart)

	r.GET("/dashboard",middleware.AdminMiddleware,controller.Dashboard)

	r.GET("/dashboard/menu",middleware.AdminMiddleware,controller.GetMenu)
	r.POST("/dashboard/menu",middleware.AdminMiddleware,controller.EditMenu)
	r.GET("/dashboard/reservation",middleware.AdminMiddleware,controller.GetReservation)
	r.POST("/dashboard/reservation",middleware.AdminMiddleware,controller.EditReservation)

	r.GET("/logout",controller.Logout)

	r.Run(":8080")
}