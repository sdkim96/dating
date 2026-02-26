package tools

import (
	"encoding/json"

	"github.com/mark3labs/mcp-go/server"
	"github.com/sdkim96/dating/internal/db"
)

func Register(s *server.MCPServer, db *db.Engine) {
	RegisterPing(s)
	RegisterCreateCard(s, db)
	RegisterListCards(s, db)
}

func ConvertMCPRequest[T any](m map[string]any) (T, error) {
	var result T
	b, err := json.Marshal(m)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(b, &result)
	return result, err
}
