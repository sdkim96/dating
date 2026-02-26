package card

import (
	"context"
	"encoding/json"

	"github.com/sdkim96/dating/internal/db"
)

type CardsCreateRequest struct {
	Cards []CardCreate `json:"cards" jsonschema:"description=The list of profile cards to be created"`
}
type CardCreate struct {
	Human    Human    `json:"human" jsonschema:"required,description=The identity of the person"`
	Position string   `json:"position,omitempty" jsonschema:"description=The job position of the person"`
	Company  *Company `json:"company,omitempty" jsonschema:"description=The company associated with the person"`
}

func Create(
	ctx context.Context,
	engine *db.Engine,
	req CardsCreateRequest,
) (int, error) {
	var entries []db.WriteEntry
	for _, card := range req.Cards {
		key := createKey(tenant, card.Human.Name)
		valueBytes, err := json.Marshal(card)
		if err != nil {
			return 0, err
		}
		meta, err := json.Marshal(map[string]string{
			"name":     card.Human.Name,
			"company":  card.Company.Name,
			"position": card.Position,
		})
		if err != nil {
			return 0, err
		}
		entries = append(entries, db.WriteEntry{
			Key:      key,
			Value:    string(valueBytes),
			Metadata: string(meta),
		})
	}

	err := engine.WriteBatch(ctx, tenant, entries)
	if err != nil {
		return 0, err
	}

	return len(req.Cards), nil
}
