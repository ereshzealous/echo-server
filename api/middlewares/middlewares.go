package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func SetAdminMiddlewares(g *echo.Group) {
	//this logs server interaction
	file, err := os.OpenFile("info-admin.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {

	}
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}]` + "\n",
		Output: file,
	}))
	//Basic Authentication
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))
}

func SetUserMiddleWares(g *echo.Group) {
	//file, err := os.OpenFile("info-user.log", os.O_CREATE|os.O_APPEND, 0644)
	/* if err != nil {

	}
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human} Error-${error}]` + "\n",
		Output: file,
	})) */
	//Basic Authentication
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "user" && password == "user123$" {
			return true, nil
		}
		return false, nil
	}))
}

func SetMainMiddleWares(e *echo.Echo) {
	e.Use(serverHeader)
}

//Custom Middleware
// ServerHeader middleware adds a `Server` header to the response.
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Custom-Header", "blah!!!")
		return next(c)
	}
}
