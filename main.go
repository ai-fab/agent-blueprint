package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"

	"your-module-name/config"
	"your-module-name/handlers"
	"your-module-name/middleware"
)

func main() {
	app := pocketbase.New()

	// Initialize PocketBase
	if err := config.InitializePocketBase(app); err != nil {
		log.Fatalf("Failed to initialize PocketBase: %v", err)
	}

	// Setup Echo routes
	e := echo.New()
	e.Use(middleware.ClientAuth(app))

	// Register routes
	handlers.RegisterRoutes(e, app)

	// Start the server
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", echo.WrapHandler(e.Router), middleware.ClientAuth(app))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
