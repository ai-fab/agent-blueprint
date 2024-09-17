package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	"service-blueprint/config"
	"service-blueprint/handlers"
	"service-blueprint/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	app := pocketbase.New()

	// Register migrate command
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Dir:         "./migrations",
		Automigrate: true,
	})

	// Initialize PocketBase (this will run migrations)
	if err := config.InitializePocketBase(app); err != nil {
		log.Fatalf("Failed to initialize PocketBase: %v", err)
	}

	// Create admin user after migrations have run
	if err := createAdminUser(app); err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
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
		e.Router.GET("/*", echo.WrapHandler(echoApp))

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func createAdminUser(app *pocketbase.PocketBase) error {
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminEmail == "" || adminPassword == "" {
		return fmt.Errorf("ADMIN_EMAIL and ADMIN_PASSWORD must be set in .env file")
	}

	admin, err := app.Dao().FindAdminByEmail(adminEmail)
	if err == nil && admin != nil {
		// Admin already exists
		return nil
	}

	admin = &models.Admin{}
	admin.Email = adminEmail
	admin.SetPassword(adminPassword)

	return app.Dao().SaveAdmin(admin)
}
