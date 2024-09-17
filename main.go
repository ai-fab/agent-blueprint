package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"log"
	"os"

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

	// Initialize PocketBase and run migrations
	if err := config.InitializePocketBase(app); err != nil {
		log.Fatalf("Failed to initialize PocketBase: %v", err)
	}

	// Ensure migrations are applied
	if err := app.Bootstrap(); err != nil {
		log.Fatalf("Failed to bootstrap PocketBase: %v", err)
	}

	// Run migrations
	if err := app.ResetBootstrapState(); err != nil {
		log.Fatalf("Failed to reset bootstrap state: %v", err)
	}

	// Create admin user after migrations have run
	if err := createAdminUser(app); err != nil {
		log.Printf("Failed to create admin user: %v", err)
		// Continue execution even if admin user creation fails
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

	if app.Dao() == nil {
		return fmt.Errorf("app.Dao() is nil, PocketBase may not be properly initialized")
	}

	admin, err := app.Dao().FindAdminByEmail(adminEmail)
	if err != nil {
		// If the error is not "record not found", return the error
		if err.Error() != "sql: no rows in result set" {
			return fmt.Errorf("error finding admin: %v", err)
		}
	}

	if admin != nil {
		log.Println("Admin user already exists")
		return nil
	}

	admin = &models.Admin{}
	admin.Email = adminEmail
	admin.SetPassword(adminPassword)

	err = app.Dao().SaveAdmin(admin)
	if err != nil {
		return fmt.Errorf("failed to save admin: %v", err)
	}

	log.Println("Admin user created successfully")
	return nil
}
