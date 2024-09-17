package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func Init_1718706525_add_login_alert_column(db dbx.Builder) error {
	dao := daos.New(db);

	collection, err := dao.FindCollectionByNameOrId("users")
	if err != nil {
		return err
	}

	// add
	new_login_alert := &models.SchemaField{}
	json.Unmarshal([]byte(`{
		"system": false,
		"id": "login_alert",
		"name": "login_alert",
		"type": "bool",
		"required": false,
		"unique": false,
		"options": {}
	}`), new_login_alert)
	collection.Schema.AddField(new_login_alert)

	return dao.SaveCollection(collection)
}
