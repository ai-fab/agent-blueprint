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

	// Manually register all migration files
	m.MustRegister(
		migrations.Init_1_initial_schema,
		migrations.Init_2_initial_admin,
		migrations.Init_1718706525_add_login_alert_column,
	)

	// Migrations will be automatically run by PocketBase
	return nil
}
