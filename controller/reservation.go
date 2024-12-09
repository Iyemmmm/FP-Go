package controller

import (
	"fmt"
	"net/http"
	"resto/initializers"
	"resto/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ViewReservation(c *gin.Context) {
	c.HTML(http.StatusOK, "reservation.html", nil)
}

func AddReservation(c *gin.Context) {
	usercookies, _ := c.Get("User")
	user := usercookies.(model.User)
	id_user := user.ID
	fmt.Println(id_user)

	nama := c.PostForm("nama")
	email := c.PostForm("email")
	note := c.PostForm("note")
	hp := c.PostForm("phone")
	dateStr := c.PostForm("date")
	parsedDate, _ := time.Parse("2006-01-02", dateStr)
	waktuStr := dateStr
	waktuStr += " "
	waktuStr += c.PostForm("time")
	waktuStr += (":00")
	parsedWaktu, _ := time.Parse("2006-01-02 15:04:05", waktuStr)
	fmt.Println(dateStr, parsedDate)
	fmt.Println(waktuStr, parsedWaktu)
	guest, _ := strconv.ParseInt(c.PostForm("guests"), 10, 64)

	reservation := model.Reservation{
		ID_User: id_user,
		Nama:    nama,
		Email:   email,
		Note:    note,
		NoHP:    hp,
		Date:    parsedDate,
		Guest:   int(guest),
		Waktu:   parsedWaktu,
	}
	result := initializers.DB.Create(&reservation)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/")
}
