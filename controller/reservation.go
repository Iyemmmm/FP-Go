package controller

import (
	"fmt"
	"log"
	"net/http"
	"resto/initializers"
	"resto/model"
	"strconv"

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
	waktuStr := c.PostForm("time")
	// parsedWaktu, _ := time.Parse("2006-01-02 15:04:05", waktuStr)
	// fmt.Println(dateStr, parsedDate)
	// fmt.Println(waktuStr, parsedWaktu)
	guest, _ := strconv.ParseInt(c.PostForm("guests"), 10, 64)

	reservation := model.Reservation{
		ID_User: id_user,
		Nama:    nama,
		Email:   email,
		Note:    note,
		NoHP:    hp,
		Date:    dateStr,
		Guest:   int(guest),
		Waktu:   waktuStr,
		Status:  "PENDING",
	}
	result := initializers.DB.Create(&reservation)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/reservation-status")
}

func ViewStatus(c *gin.Context) {
	usercookies, _ := c.Get("User")
	user := usercookies.(model.User)
	var reservation []model.Reservation
	result := initializers.DB.Preload("User").Where("id_user = ?", user.ID).Find(&reservation)

	if result.Error != nil {
		log.Println("Error fetching resrvation:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
		return
	}
	log.Println("Reservation data:", reservation)
	// Kirim data ke template HTML
	c.HTML(http.StatusOK, "reservation-status.html", gin.H{"reservations": reservation})

}
