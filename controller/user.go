package controller

import (
	"net/http"
	"os"
	"resto/initializers"
	"resto/model"
	"resto/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ViewRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func CreateUser(c *gin.Context) {
	nama := c.PostForm("nama")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmpassword := c.PostForm("confirmpassword")
	hp := c.PostForm("nohp")
	role := "user"

	// Validasi password
	if password != confirmpassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password and Confirm Password do not match"})
		return
	}

	if len(password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Minimum Password is 8 characters"})
		return
	}
	var user model.User
	query := initializers.DB.Where("email = ?", email).First(&user)
	if query != nil && query.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email already in use."})
		return
	}

	password, _ = util.GenereateHasedPassword(password)

	account := model.User{
		Nama:     nama,
		Email:    email,
		Password: password,
		NoHP:     hp,
		Role:     role,
	}
	result := initializers.DB.Create(&account)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func ViewLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginUser(c *gin.Context) {
	godotenv.Load()
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user model.User
	if result := initializers.DB.Where("email = ?", email).First(&user).Error; result != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	// fmt.Println(user.Password)
	err := util.CompareHashedPassword(user.Password, password)
	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, tokenerr := token.SignedString([]byte(os.Getenv("SECRET")))

	if tokenerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid to create token",
		})
		return
	}
	if user.Role == "admin" {
		c.Redirect(http.StatusFound, "/dashboard")
	} else {
		c.Redirect(http.StatusFound, "/")
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.Redirect(http.StatusFound, "/login")
}
