package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/BEkasss11/golang/initializers"
	"github.com/BEkasss11/golang/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body models.SignUpInput

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	//Create User
	user := models.User{
		Email:    body.Email,
		Password: string(hash),
		Username: body.UserName,
		Role:     "User",
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create User.",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{"message": "User created!"})
}

func SignIn(c *gin.Context) {
	var body models.SignInInput

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "username = ?", body.UserName)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid phone number.",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user.ID,
		"userRole": user.Role,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token.",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func SignOut(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Signed out."})
}

func ValidateUser(c *gin.Context) {
	user, err := c.Get("user")

	if !err {
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetPayloadFromToken(c *gin.Context) gin.H {
	tokenCookie, err := c.Request.Cookie("Authorization")
	if err != nil {
		return nil
	}
	tokenString := tokenCookie.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {

		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		response := gin.H{
			"userID":   claims["userID"],
			"userRole": claims["userRole"],
		}
		return response
	} else {
		return nil
	}
}
