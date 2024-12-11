package controller

import (
	"fmt"
	"log"
	"net/http"
	"resto/initializers"
	"resto/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ViewCart(c *gin.Context) {
	usercookies, _ := c.Get("User")
	user := usercookies.(model.User)
	var cart []model.Cart
	carts := initializers.DB.Preload("Menu").Where("id_user = ?", user.ID).Find(&cart)

	if carts.Error != nil {
		log.Println("Error fetching carts:", carts.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch carts"})
		return
	}

	log.Println("Cart data:", cart)
	// Kirim data ke template HTML
	c.HTML(http.StatusOK, "cart.html", gin.H{"carts": cart})
}

func AddCart(c *gin.Context) {
	usercookies, _ := c.Get("User")
	user := usercookies.(model.User)
	id_user := user.ID
	id_menu, _ := strconv.ParseUint(c.PostForm("id_menu"), 10, 64)
	fmt.Println(id_user, id_menu)
	quantity, _ := strconv.ParseFloat(c.PostForm("quantity"), 64)
	notes := c.PostForm("note")
	fmt.Println(notes)
	var menu model.Menu
	query := initializers.DB.Where("id = ?", id_menu).First(&menu)
	if query.Error != nil && query.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menu not Found."})
	}
	total := menu.Harga * quantity
	cart := model.Cart{
		ID_User:  uint(id_user),
		ID_Menu:  uint(id_menu),
		Status:   "PENDING",
		Note:     notes,
		Total:    total,
		Quantity: int(quantity),
	}
	result := initializers.DB.Create(&cart)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/cart")
}
