package config

import (
	"fmt"

	"github.com/pocketbase/pocketbase"
)

func InitializePocketBase(app *pocketbase.PocketBase) error {
	// Initialize the database
	if err := app.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap PocketBase: %w", err)
	}

	// Migrations are handled automatically by migratecmd
	// No need to run migrations manually

	return nil
}
