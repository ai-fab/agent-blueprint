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

	// Ensure the database is up to date
	if err := app.ResetBootstrapState(); err != nil {
		return fmt.Errorf("failed to reset bootstrap state: %w", err)
	}

	return nil
}
