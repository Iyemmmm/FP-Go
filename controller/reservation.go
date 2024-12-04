package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ViewReservation(c *gin.Context) {
	c.HTML(http.StatusOK, "reservation.html", nil)
}