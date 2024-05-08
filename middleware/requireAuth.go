package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/BEkasss11/golang/enums"
	"github.com/BEkasss11/golang/initializers"
	"github.com/BEkasss11/golang/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {

	// get the cookie off req
	tokenString, err := c.Cookie("Authorization")

	if err != nil || tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized."})
		return
	}

	//decode
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.DB.First(&user, claims["userID"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Set("role", user.Role)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func IsAdmin(c *gin.Context) {
	role, exists := c.Get("role")
	fmt.Println(role)
	fmt.Println(exists)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not admin."})
		c.Abort()
		return
	}
	if role != "Admin" {
		fmt.Println(role)
		fmt.Println(enums.ADMIN)
		fmt.Println(role != enums.ADMIN)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not admin."})
		c.Abort()
		return
	}
	c.Next()
}
