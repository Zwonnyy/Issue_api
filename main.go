// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"issueapi/router"
)

func main() {
	r := gin.Default()
	router.InitRoutes(r)
	r.Run(":8080")
}
