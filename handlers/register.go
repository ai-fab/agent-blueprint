package handlers

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

func RegisterRoutes(e *echo.Echo, app *pocketbase.PocketBase) {
	e.POST("/api/projects", createProject(app))
	e.GET("/api/projects", listProjects(app))
	e.GET("/api/projects/:id/status", getProjectStatus(app))
}
