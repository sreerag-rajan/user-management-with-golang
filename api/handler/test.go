package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	fmt.Println("Hitting the test route")
	c.JSON(http.StatusOK, gin.H{"message": "Test successfull"})

}
