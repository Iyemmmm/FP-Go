package controller

import (
	"fmt"
	"net/http"
	// "resto/model"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	is_login:= c.GetBool("id")

	fmt.Println(is_login)
	// Kirim data ke template
	c.HTML(http.StatusOK, "index.html", gin.H{
		"id": is_login,
	})
}


