package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

func RegisterRoutes(e *echo.Echo, app *pocketbase.PocketBase) {
	e.POST("/projects", createProject(app))
	e.GET("/projects", listProjects(app))
	e.GET("/projects/:id/status", getProjectStatus(app))
}

func createProject(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientID := c.Request().Header.Get("X-Client-ID")

		record := models.NewRecord(app.Dao().FindCollectionByNameOrId("projects"))
		form := forms.NewRecordUpsert(app, record)

		if err := c.Bind(form); err != nil {
			return err
		}

		form.SetDataValue("client_id", clientID)

		if err := form.Submit(); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, record)
	}
}

func listProjects(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientID := c.Request().Header.Get("X-Client-ID")

		records, err := app.Dao().FindRecordsByExpr("projects", "client_id = {:clientID}", map[string]interface{}{"clientID": clientID})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, records)
	}
}

func getProjectStatus(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID := c.Param("id")

		record, err := app.Dao().FindRecordById("projects", projectID)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Project not found")
		}

		return c.JSON(http.StatusOK, map[string]string{"status": record.GetString("status")})
	}
}
