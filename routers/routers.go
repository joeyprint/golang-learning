package routers

import (
	"live/authentication/src/packages/user"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	router *gin.Engine
}

func Router() Routes {
	router := Routes{router: gin.Default()}

	api := router.router.Group("/api")

	user.AddUserController(api)

	return router
}

func (routes Routes) Run(addr ...string) error {
	return routes.router.Run()
}
