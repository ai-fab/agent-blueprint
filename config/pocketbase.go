package config

import (
	"fmt"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

func InitializePocketBase(app *pocketbase.PocketBase) error {
	// Initialize the database
	if err := app.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap PocketBase: %w", err)
	}

	collections := []struct {
		name   string
		schema schema.Schema
	}{
		{
			name: "settings",
			schema: schema.Schema{
				"user":  {&schema.SchemaField{Type: schema.FieldTypeText}},
				"key":   {&schema.SchemaField{Type: schema.FieldTypeText}},
				"value": {&schema.SchemaField{Type: schema.FieldTypeText}},
			},
		},
		{
			name: "client_applications",
			schema: schema.Schema{
				"client_id":     {&schema.SchemaField{Type: schema.FieldTypeText}},
				"client_secret": {&schema.SchemaField{Type: schema.FieldTypeText}},
			},
		},
		{
			name: "projects",
			schema: schema.Schema{
				"name":      {&schema.SchemaField{Type: schema.FieldTypeText}},
				"client_id": {&schema.SchemaField{Type: schema.FieldTypeText}},
				"status":    {&schema.SchemaField{Type: schema.FieldTypeText}},
			},
		},
	}

	for _, col := range collections {
		collection, err := app.Dao().FindCollectionByNameOrId(col.name)
		if err != nil {
			// Collection doesn't exist, create it
			collection = &models.Collection{
				Name:   col.name,
				Type:   models.CollectionTypeBase,
				Schema: col.schema,
			}
			if err := app.Dao().SaveCollection(collection); err != nil {
				return fmt.Errorf("failed to create collection %s: %w", col.name, err)
			}
		} else {
			// Collection exists, update its schema
			collection.Schema = col.schema
			if err := app.Dao().SaveCollection(collection); err != nil {
				return fmt.Errorf("failed to update collection %s: %w", col.name, err)
			}
		}
	}

	return nil
}

// Helper function to create a schema
func createSchema(fields []*schema.SchemaField) schema.Schema {
	return schema.NewSchema(fields...)
}
