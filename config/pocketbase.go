package config

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

func InitializePocketBase(app *pocketbase.PocketBase) error {
	// Create settings collection
	settingsCollection := &models.Collection{
		Name: "settings",
		Type: models.CollectionTypeBase,
		Schema: createSchema([]*schema.SchemaField{
			{Name: "user", Type: schema.FieldTypeText},
			{Name: "key", Type: schema.FieldTypeText},
			{Name: "value", Type: schema.FieldTypeText},
		}),
	}

	// Create client_applications collection
	clientAppsCollection := &models.Collection{
		Name: "client_applications",
		Type: models.CollectionTypeBase,
		Schema: createSchema([]*schema.SchemaField{
			{Name: "client_id", Type: schema.FieldTypeText},
			{Name: "client_secret", Type: schema.FieldTypeText},
		}),
	}

	// Create projects collection
	projectsCollection := &models.Collection{
		Name: "projects",
		Type: models.CollectionTypeBase,
		Schema: createSchema([]*schema.SchemaField{
			{Name: "name", Type: schema.FieldTypeText},
			{Name: "client_id", Type: schema.FieldTypeText},
			{Name: "status", Type: schema.FieldTypeText},
		}),
	}

	collections := []*models.Collection{settingsCollection, clientAppsCollection, projectsCollection}

	for _, collection := range collections {
		if err := app.Dao().SaveCollection(collection); err != nil {
			return err
		}
	}

	return nil
}

// Helper function to create a schema
func createSchema(fields []*schema.SchemaField) schema.Schema {
	return schema.NewSchema(fields...)
}
