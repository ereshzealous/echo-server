package router

import (
	"echoserver/api"
	"echoserver/api/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//create groups
	adminGroup := e.Group("/admin")
	userGroup := e.Group("/user")

	//set all middlewares
	middlewares.SetMainMiddleWares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetUserMiddleWares(userGroup)

	//set main routes
	api.MainGroup(userGroup)

	//set groupRoutes
	api.AdminGroup(adminGroup)

	return e
}
