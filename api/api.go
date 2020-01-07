package api

import (
	"echoserver/handlers"
	"github.com/labstack/echo"
)

func MainGroup(g *echo.Group) {
	// Route / to handler function
	g.GET("/health-check", handlers.HealthCheck)

	g.GET("/cats/:data", handlers.GetCats)
	g.POST("/cats", handlers.AddCat)

}

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
