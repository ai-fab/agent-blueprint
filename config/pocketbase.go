package config

import (
	"fmt"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/migrate"
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

	// Run migrations
	if err := migrate.Up(app.DB(), migrationsDir); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
