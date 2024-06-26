package final_project

import (
	"fmt"
	"github.com/BEkasss11/golang/middleware"
	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) bool {
	userRole := middleware.GetPayloadFromToken(c)["userRole"]
	fmt.Println(userRole)
	if userRole != "Admin" {
		return false
	}
	return true
}

func IsAuthorizedOrReadOnly(c *gin.Context) bool {
	loggedUser := middleware.GetPayloadFromToken(c)
	if loggedUser == nil {
		return false
	}
	return true
}
