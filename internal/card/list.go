package card

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sdkim96/dating/internal/db"
)

type CardQuery struct {
	Field string `json:"field,omitempty" jsonschema:"description=The field to search. Only 'name'&#44; 'company'&#44; 'position' is allowed"`
	Value string `json:"value,omitempty" jsonschema:"description=The value to search for in the specified field"`
}
type CardsListRequest struct {
	Size    int         `json:"size,omitempty" jsonschema:"description=The number of items per page for pagination"`
	Offset  int         `json:"offset,omitempty" jsonschema:"description=The offset for pagination"`
	Queries []CardQuery `json:"queries,omitempty" jsonschema:"description=The list of search queries to filter the profile cards"`
}

type CardsListResponse struct {
	Cards []Card `json:"cards" jsonschema:"description=The list of profile cards matching the search criteria"`
}

func List(
	ctx context.Context,
	engine *db.Engine,
	req CardsListRequest,
) (CardsListResponse, error) {
	var metadataFilters []db.MetadataFilter
	for _, q := range req.Queries {
		if q.Field != "name" && q.Field != "company" && q.Field != "position" {
			return CardsListResponse{}, fmt.Errorf("invalid query field: %s", q.Field)
		}
		metadataFilters = append(metadataFilters, db.MetadataFilter{
			Field: q.Field,
			Value: q.Value,
		})
	}
	opt := &db.ReadAllOption{
		OrderBy:  "created_at",
		Limit:    req.Size,
		Offset:   req.Offset,
		Metadata: metadataFilters,
	}

	raw, err := engine.ReadAll(ctx, tenant, opt)
	if err != nil {
		return CardsListResponse{}, err
	}
	if raw == "" {
		return CardsListResponse{Cards: []Card{}}, nil
	}

	var result map[string]map[string]string
	if err := json.Unmarshal([]byte(raw), &result); err != nil {
		return CardsListResponse{}, err
	}

	var cards []Card
	for _, v := range result[tenant] {
		var c Card
		if err := json.Unmarshal([]byte(v), &c); err != nil {
			continue
		}
		cards = append(cards, c)
	}

	return CardsListResponse{Cards: cards}, nil
}
