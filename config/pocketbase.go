package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"plugin"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	m "github.com/pocketbase/pocketbase/migrations"
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

	// Automatically discover and register migration files
	if err := registerMigrations(migrationsDir); err != nil {
		return fmt.Errorf("failed to register migrations: %w", err)
	}

	// Migrations will be automatically run by PocketBase
	return nil
}

func registerMigrations(migrationsDir string) error {
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".go" {
			p, err := plugin.Open(filepath.Join(migrationsDir, file.Name()))
			if err != nil {
				return fmt.Errorf("failed to open migration file %s: %w", file.Name(), err)
			}

			initFunc, err := p.Lookup("init")
			if err != nil {
				return fmt.Errorf("failed to find init function in %s: %w", file.Name(), err)
			}

			initFunc.(func())()
		}
	}

	return nil
}
