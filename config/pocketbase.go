package config

import (
	"fmt"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func InitializePocketBase(app *pocketbase.PocketBase) error {
	// Initialize the database
	if err := app.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap PocketBase: %w", err)
	}

	// Load and execute migrations
	migrationsDir := "./migrations"
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Dir:         migrationsDir,
		Automigrate: true,
	})

	// Run migrations manually
	if err := runMigrations(app); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func runMigrations(app *pocketbase.PocketBase) error {
	collection, err := app.Dao().FindCollectionByNameOrId("_migrations")
	if err != nil {
		return fmt.Errorf("failed to find migrations collection: %w", err)
	}

	migrations, err := migratecmd.NewMigrationsRunner(app.Dao(), collection)
	if err != nil {
		return fmt.Errorf("failed to create migrations runner: %w", err)
	}

	return migrations.Run(migratecmd.RunOptions{})
}
