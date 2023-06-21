package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sreerag-rajan/user-management-with-go/controller"
)

func Routes(r *gin.Engine) {

	r.GET("/user/test", controller.Test)
}
