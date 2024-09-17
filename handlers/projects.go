package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

func createProject(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientID := c.Request().Header.Get("X-Client-ID")

		collection, err := app.Dao().FindCollectionByNameOrId("projects")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find projects collection")
		}

		record := models.NewRecord(collection)
		if err := json.NewDecoder(c.Request().Body).Decode(record); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input data")
		}

		// Validate required fields
		if record.Get("name") == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Project name is required")
		}

		record.Set("client_id", clientID)

		if err := app.Dao().SaveRecord(record); err != nil {
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

		query := app.Dao().RecordQuery("projects").
			Where(dbx.HashExp{"client_id": clientID}).
			Limit(int64(perPage)).
			Offset(int64((page - 1) * perPage))

		records := []models.Record{}
		err := query.All(&records)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch projects: "+err.Error())
		}

		totalRecords := 0
		err = app.Dao().RecordQuery("projects").
			Where(dbx.HashExp{"client_id": clientID}).
			Select("count(*)").
			Row(&totalRecords)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to count projects: "+err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"items": records,
			"page":  page,
			"total": totalRecords,
		})
	}
}

func getProjectStatus(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID := c.PathParam("id")
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
