package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2820579054")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"updateRule": "owner.id = @request.auth.id &&\n// only allow updating fields: isFavorite\n@request.body.subject = null &&\n@request.body.body = null &&\n@request.body.headers = null &&\n@request.body.incoming = null &&\n@request.body.outgoing = null &&\n@request.body.sender = null &&\n@request.body.recipient = null &&\n@request.body.created = null &&\n@request.body.updated = null"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"hidden": false,
			"id": "bool4089940794",
			"name": "isFavorite",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "bool"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2820579054")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"updateRule": null
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("bool4089940794")

		return app.Save(collection)
	})
}
