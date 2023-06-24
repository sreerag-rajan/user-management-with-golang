package auth

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sreerag-rajan/user-management-with-go/model"
)

func Register(c *gin.Context) {
	var registerBody model.User

	err := c.ShouldBindJSON(&registerBody)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PAYLAOD: %#v", registerBody)

}
