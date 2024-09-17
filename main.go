package main

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"

	"service-blueprint/config"
	"service-blueprint/handlers"
	"service-blueprint/middleware"
)

func main() {
	app := pocketbase.New()

	// Initialize PocketBase
	if err := config.InitializePocketBase(app); err != nil {
		log.Fatalf("Failed to initialize PocketBase: %v", err)
	}

	// Start the server
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Create a new Echo instance
		echoApp := echo.New()

		// Apply middleware
		echoApp.Use(middleware.ClientAuth(app))

		// Register routes
		handlers.RegisterRoutes(echoApp, app)

		// Mount Echo to the root
		e.Router.GET("/*", func(c echo.Context) error {
			echoApp.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
