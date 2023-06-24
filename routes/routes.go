package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sreerag-rajan/user-management-with-go/api/handler"
	"github.com/sreerag-rajan/user-management-with-go/api/handler/auth"
)

func Routes(r *gin.Engine) {

	r.GET("/user/test", handler.Test)
	r.POST("/auth/register", auth.Register)
}
