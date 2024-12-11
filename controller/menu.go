package controller

import (
	"log"
	"net/http"
	"path/filepath"
	"resto/initializers"
	"resto/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ViewMenu(c *gin.Context) {
	var menu []model.Menu
	products := initializers.DB.Find(&menu)

	if products.Error != nil {
		log.Println("Error fetching products:", products.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	log.Println("Menu data:", menu)
	// Kirim data ke template HTML
	c.HTML(http.StatusOK, "menu.html", gin.H{
		"products": menu,
	})
}

func AddMenu(c *gin.Context, uploadPath string) {
	nama := c.PostForm("nama")
	hargastr := c.PostForm("harga")
	harga, _ := strconv.ParseFloat(hargastr, 64)
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}
	fileName := filepath.Base(file.Filename)
	filePath := filepath.Join(uploadPath, fileName)
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		log.Println("Error saving file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	menu := model.Menu{
		Nama:  nama,
		Harga: harga,
		Img:   fileName,
	}
	result := initializers.DB.Create(&menu)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/menu")
}
