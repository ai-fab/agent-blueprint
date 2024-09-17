package config

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

func InitializePocketBase(app *pocketbase.PocketBase) error {
	// Create settings collection
	settingsCollection := &models.Collection{
		Name: "settings",
		Type: models.CollectionTypeBase,
		Schema: models.Schema{
			&models.SchemaField{Name: "user", Type: "text"},
			&models.SchemaField{Name: "key", Type: "text"},
			&models.SchemaField{Name: "value", Type: "text"},
		},
	}

	// Create client_applications collection
	clientAppsCollection := &models.Collection{
		Name: "client_applications",
		Type: models.CollectionTypeBase,
		Schema: models.Schema{
			&models.SchemaField{Name: "client_id", Type: "text"},
			&models.SchemaField{Name: "client_secret", Type: "text"},
		},
	}

	// Create projects collection
	projectsCollection := &models.Collection{
		Name: "projects",
		Type: models.CollectionTypeBase,
		Schema: models.Schema{
			&models.SchemaField{Name: "name", Type: "text"},
			&models.SchemaField{Name: "client_id", Type: "text"},
			&models.SchemaField{Name: "status", Type: "text"},
		},
	}

	collections := []*models.Collection{settingsCollection, clientAppsCollection, projectsCollection}

	for _, collection := range collections {
		if err := app.Dao().SaveCollection(collection); err != nil {
			return err
		}
	}

	return nil
}
