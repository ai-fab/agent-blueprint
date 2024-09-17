package middleware

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

func ClientAuth(app *pocketbase.PocketBase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			clientID := c.Request().Header.Get("X-Client-ID")
			clientSecret := c.Request().Header.Get("X-Client-Secret")

			if clientID == "" || clientSecret == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing client credentials")
			}

			record, err := app.Dao().FindFirstRecordByData("client_applications", "client_id", clientID)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid client credentials")
			}

			if record.GetString("client_secret") != clientSecret {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid client credentials")
			}

			return next(c)
		}
	}
}
