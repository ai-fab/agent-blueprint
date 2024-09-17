package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	pbmodels "github.com/pocketbase/pocketbase/models"
)

func RegisterRoutes(e *echo.Echo, app *pocketbase.PocketBase) {
	e.POST("/projects", createProject(app))
	e.GET("/projects", listProjects(app))
	e.GET("/projects/:id/status", getProjectStatus(app))
}

func createProject(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientID := c.Request().Header.Get("X-Client-ID")

		record := pbmodels.NewRecord(app.Dao().FindCollectionByNameOrId("projects"))
		form := forms.NewRecordUpsert(app, record)

		if err := c.Bind(form); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input data")
		}

		// Validate required fields
		if form.GetString("name") == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Project name is required")
		}

		form.SetDataValue("client_id", clientID)

		if err := form.Submit(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create project: "+err.Error())
		}

		return c.JSON(http.StatusCreated, record)
	}
}

func listProjects(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientID := c.Request().Header.Get("X-Client-ID")

		page, _ := strconv.Atoi(c.QueryParam("page"))
		if page < 1 {
			page = 1
		}
		perPage, _ := strconv.Atoi(c.QueryParam("per_page"))
		if perPage < 1 || perPage > 100 {
			perPage = 20
		}

		records, err := app.Dao().FindRecordsByExpr("projects",
			"client_id = {:clientID}",
			map[string]interface{}{"clientID": clientID},
			perPage,
			(page-1)*perPage,
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch projects: "+err.Error())
		}

		totalRecords, err := app.Dao().FindRecordsByExpr("projects", "client_id = {:clientID}", map[string]interface{}{"clientID": clientID})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to count projects: "+err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"items": records,
			"page":  page,
			"total": len(totalRecords),
		})
	}
}

func getProjectStatus(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID := c.Param("id")
		clientID := c.Request().Header.Get("X-Client-ID")

		record, err := app.Dao().FindRecordById("projects", projectID)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Project not found")
		}

		if record.GetString("client_id") != clientID {
			return echo.NewHTTPError(http.StatusForbidden, "Access denied")
		}

		return c.JSON(http.StatusOK, map[string]string{"status": record.GetString("status")})
	}
}
