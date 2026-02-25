package tools

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	"github.com/sdkim96/dating/internal/card"
)

func Register(s *server.MCPServer) {
	RegisterPing(s)
	RegisterCreateCard(s)
	RegisterListCards(s)
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

func HashCardID(card card.CardCreate) string {
	// For simplicity, we can use a hash of the name and position as the card ID.
	// In a real application, you might want to use a more robust method for generating unique IDs.
	return fmt.Sprintf("%x", sha256.Sum256([]byte(card.Human.Name+card.Position)))
}
