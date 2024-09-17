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

		record := models.NewRecord(collection)
		record.Set("client_id", "test_client")
		record.Set("client_secret", "test_secret")

		return dao.SaveRecord(record)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("client_applications")
		if err != nil {
			return err
		}

		record, err := dao.FindFirstRecordByData(collection.Id, "client_id", "test_client")
		if err != nil {
			return nil // If the record doesn't exist, we don't need to delete it
		}

		return dao.DeleteRecord(record)
	})
}
