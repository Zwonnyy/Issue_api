// router/router.go
package router

import (
	"issueapi/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/issue", controllers.CreateIssue)
	r.GET("/issues", controllers.GetIssues)
	r.GET("/issue/:id", controllers.GetIssueByID)
	r.PATCH("/issue/:id", controllers.UpdateIssue)
}
