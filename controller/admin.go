package controller

import (
	"net/http"
	"resto/initializers"
	"resto/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}

func GetMenu(c *gin.Context) {
	var cart []model.Cart
	result := initializers.DB.Preload("User").Preload("Menu").Find(&cart)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
	}

	c.HTML(http.StatusOK, "adminMenu.html", gin.H{"cart": cart})
}
func EditMenu(c *gin.Context) {
	idStr, _ := strconv.ParseUint(c.PostForm("idmenu"), 10, 64)
	action := c.PostForm("action")
	id := uint(idStr)
	var cart model.Cart
	result := initializers.DB.Where("id = ?", id).Find(&cart)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch carts"})
		return
	}
	if action == "Edit" {
		cart.Status = "COMPLETED"
		initializers.DB.Save(&cart)
	}else if action == "Delete"{
		initializers.DB.Delete(&cart)
	}
	var carts []model.Cart
	if err := initializers.DB.Preload("User").Preload("Menu").Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load carts"})
		return
	}
	c.HTML(http.StatusOK, "adminMenu.html", gin.H{"cart": carts})
}

func GetReservation(c *gin.Context) {
	var reservation []model.Reservation
	result := initializers.DB.Preload("User").Find(&reservation)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
	}

	c.HTML(http.StatusOK, "adminReservation.html", gin.H{"reservations": reservation})
}
func EditReservation(c *gin.Context) {
	idStr, _ := strconv.ParseUint(c.PostForm("idreservation"), 10, 64)
	action := c.PostForm("action")
	id := uint(idStr)
	var reservation model.Reservation
	result := initializers.DB.Where("id = ?", id).Find(&reservation)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
		return
	}
	if action == "Edit" {
		reservation.Status = "COMPLETED"
		initializers.DB.Save(&reservation)
	}else if action == "Delete"{
		initializers.DB.Delete(&reservation)
	}
	var reservations []model.Reservation
	if err := initializers.DB.Preload("User").Find(&reservations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load reservations"})
		return
	}
	c.HTML(http.StatusOK, "adminReservation.html", gin.H{"reservations": reservations})
}
