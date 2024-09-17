package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
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

	// Run migrations and create admin user
	if err := app.Bootstrap(); err != nil {
		log.Fatalf("Failed to bootstrap: %v", err)
	}

	if err := runMigrations(app); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	if err := createAdminUser(app); err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

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
		e.Router.GET("/*", echo.WrapHandler(echoApp))

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func runMigrations(app *pocketbase.PocketBase) error {
	return app.Dao().RunMigrations()
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
