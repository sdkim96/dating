package card

type CardListRequest struct {
	Query  string `json:"query,omitempty" jsonschema:"description=The search query to filter cards by name or company"`
	Size   int    `json:"page,omitempty" jsonschema:"description=The page number for pagination"`
	Offset int    `json:"size,omitempty" jsonschema:"description=The number of items per page for pagination"`
}

type CardListResponse struct {
	Cards []Card `json:"cards" jsonschema:"description=The list of profile cards matching the search criteria"`
}
