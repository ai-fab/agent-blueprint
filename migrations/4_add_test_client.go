package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("client_applications")
		if err != nil {
			return err
		}

		clients := []struct {
			clientID     string
			clientSecret string
		}{
			{"test_client_1", "test_secret_1"},
			{"test_client_2", "test_secret_2"},
		}

		for _, client := range clients {
			record := models.NewRecord(collection)
			record.Set("client_id", client.clientID)
			record.Set("client_secret", client.clientSecret)

			if err := dao.SaveRecord(record); err != nil {
				return err
			}
		}

		return nil
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("client_applications")
		if err != nil {
			return err
		}

		clients := []string{"test_client_1", "test_client_2"}

		for _, clientID := range clients {
			record, err := dao.FindFirstRecordByData(collection.Id, "client_id", clientID)
			if err != nil {
				continue // If the record doesn't exist, we don't need to delete it
			}

			if err := dao.DeleteRecord(record); err != nil {
				return err
			}
		}

		return nil
	})
}
