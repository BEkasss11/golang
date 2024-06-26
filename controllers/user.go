package controllers

import (
	"net/http"
	final_project "github.com/BEkasss11/golang"
	"github.com/BEkasss11/golang/initializers"
	"github.com/BEkasss11/golang/models"
	"github.com/gin-gonic/gin"
)

func ListOfUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Unscoped().Table("users").Find(&users)
	if len(users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't retrieve data!"})
		return
	}

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized."})
		return
	}

	isAdmin := final_project.IsAdmin(c)

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not Admin."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserByID(c *gin.Context) {
	var user models.User

	if err := initializers.DB.Where("id = ?", c.Param("id")).Unscoped().Table("users").First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	isAuth := final_project.IsAuthorizedOrReadOnly(c)

	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized."})
		return
	}

	isAdmin := final_project.IsAdmin(c)

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not Admin."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUserByID(c *gin.Context) {
	var user models.User
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	isAuth := final_project.IsAuthorizedOrReadOnly(c)
	if !isAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized."})
		return
	}

	isAdmin := final_project.IsAdmin(c)

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not Admin."})
		return
	}

	initializers.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}
