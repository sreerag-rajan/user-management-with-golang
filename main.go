package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sreerag-rajan/user-management-with-go/config/db"
	"github.com/sreerag-rajan/user-management-with-go/routes"
)

func main() {

	r := gin.Default()

	db.LoadConfig()

	//Routes being imported from the routes folder
	routes.Routes(r)

	r.Run(":8009")

}
