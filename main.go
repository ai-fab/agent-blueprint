package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"log"

	"service-blueprint/config"
	"service-blueprint/handlers"
	"service-blueprint/middleware"
	_ "service-blueprint/migrations"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	app := pocketbase.New()

	// Register migrate command
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Dir:         "migrations",
		Automigrate: false,
	})

	// Initialize PocketBase
	if err := config.InitializePocketBase(app); err != nil {
		log.Fatalf("Failed to initialize PocketBase: %v", err)
	}

	// Start the server
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		log.Println("Server is starting...")
		// Create a new Echo instance
		echoApp := echo.New()

		// Apply middleware
		echoApp.Use(middleware.ClientAuth(app))

		// Register routes
		handlers.RegisterRoutes(echoApp, app)

		// Mount Echo to the root
		e.Router.ANY("/*", echo.WrapHandler(echoApp))

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
