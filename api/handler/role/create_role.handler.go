package role

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sreerag-rajan/user-management-with-go/model"
	"github.com/sreerag-rajan/user-management-with-go/service"
)

func CreateRoleHandler(c *gin.Context) {
	fmt.Println("ENTERING HANDLER")
	var role *model.RolePayload

	err := c.BindJSON(&role)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	newRole, err := service.CreateUserRoleService(role)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "ROLE CREATED", "Role": newRole})
}
