package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/sdkim96/dating/internal/card"
	"github.com/sdkim96/dating/internal/db"
)

const CreateCardDescription = `
Create a profile card with the following fields:
`
const ListCardsDescription = `
List all profile cards.
`

func RegisterCreateCard(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool(
			"create_card",
			mcp.WithDescription(CreateCardDescription),
			mcp.WithInputSchema[card.CardCreate](),
		),
		handleCreateCard,
	)
}

func handleCreateCard(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {

	convertedCard, err := ConvertMCPRequest[card.CardCreate](req.GetArguments())
	if err != nil {
		return nil, fmt.Errorf("failed to convert arguments: %w", err)
	}

	fmt.Println("Received card creation request:", convertedCard)

	cardID := HashCardID(convertedCard)
	marshaled, err := json.Marshal(convertedCard)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal card: %w", err)
	}

	db.Store.Write(cardID, string(marshaled))
	return mcp.NewToolResultText("Card created"), nil
}
func RegisterListCards(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool(
			"list_cards",
			mcp.WithDescription(ListCardsDescription),
			mcp.WithInputSchema[card.CardListRequest](),
			mcp.WithOutputSchema[card.CardListResponse](),
		),
		handleListCards,
	)
}

func handleListCards(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {

	convertedRequest, err := ConvertMCPRequest[card.CardListRequest](req.GetArguments())
	if err != nil {
		return nil, fmt.Errorf("failed to convert arguments: %w", err)
	}

	cardEntries := db.Store.ReadAll()
	var Cards []card.Card
	for _, entry := range cardEntries {
		var c card.Card
		err := json.Unmarshal([]byte(entry.Value), &c)
		if err != nil {
			fmt.Printf("failed to unmarshal card with key %s: %v\n", entry.Key, err)
			continue
		}
		Cards = append(Cards, c)
	}

	fmt.Println("Received list cards request:", convertedRequest)
	response := card.CardListResponse{Cards: Cards}
	return mcp.NewToolResultJSON(response)
}
