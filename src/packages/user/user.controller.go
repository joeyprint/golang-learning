package user

import (
	"github.com/gin-gonic/gin"
)

func AddUserController(routerGroup *gin.RouterGroup) {
	userRoutes := routerGroup.Group("/users")

	userRoutes.GET("/", getUsers)
	userRoutes.GET("/:id", getUsersById)
}
