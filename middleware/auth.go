package middleware

import (
	"fmt"
	"net/http"
	"os"
	"resto/initializers"
	"resto/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func RequireAuth(c *gin.Context) {
	godotenv.Load()
	// GET TOKEN
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
        c.Abort()
        return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix())>claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user model.User
		initializers.DB.First(&user,claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if user.Role != "user"{
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
            c.Abort()
            return
		}

		c.Set("User",user)
		c.Next()
		// fmt.Println(claims["sub"], claims["exp"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
func AdminMiddleware(c *gin.Context) {
	godotenv.Load()
	// GET TOKEN
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
        c.Abort()
        return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix())>claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user model.User
		initializers.DB.First(&user,claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		
		if user.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
            c.Abort()
            return
		}
		c.Set("User",user)
		c.Next()
		// fmt.Println(claims["sub"], claims["exp"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}