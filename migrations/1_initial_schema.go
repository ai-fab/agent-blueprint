package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collections := []struct {
			name   string
			schema schema.Schema
		}{
			{
				name: "settings",
				schema: schema.NewSchema(
					&schema.SchemaField{
						Name: "user",
						Type: "text",
					},
					&schema.SchemaField{
						Name: "key",
						Type: "text",
					},
					&schema.SchemaField{
						Name: "value",
						Type: "text",
					},
				),
			},
			{
				name: "client_applications",
				schema: schema.NewSchema(
					&schema.SchemaField{
						Name: "client_id",
						Type: "text",
					},
					&schema.SchemaField{
						Name: "client_secret",
						Type: "text",
					},
				),
			},
			{
				name: "projects",
				schema: schema.NewSchema(
					&schema.SchemaField{
						Name: "name",
						Type: "text",
					},
					&schema.SchemaField{
						Name: "client_id",
						Type: "text",
					},
					&schema.SchemaField{
						Name: "status",
						Type: "text",
					},
				),
			},
		}

		for _, col := range collections {
			collection := &models.Collection{
				Name:   col.name,
				Type:   models.CollectionTypeBase,
				Schema: col.schema,
			}
			if err := dao.SaveCollection(collection); err != nil {
				return err
			}
		}

		return nil
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collections := []string{"settings", "client_applications", "projects"}

		for _, name := range collections {
			collection, err := dao.FindCollectionByNameOrId(name)
			if err != nil {
				return err
			}
			if err := dao.DeleteCollection(collection); err != nil {
				return err
			}
		}

		return nil
	})
}
