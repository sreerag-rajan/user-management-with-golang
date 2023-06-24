package routes

import (
	"github.com/gin-gonic/gin"
	handler_test "github.com/sreerag-rajan/user-management-with-go/api/handler"
)

func Routes(r *gin.Engine) {

	r.GET("/user/test", handler_test.Test)
}
