package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/sdkim96/dating/internal/card"
	"github.com/sdkim96/dating/internal/db"
)

const cardsCreateDesc = `
Create a profile card with the following fields:
`
const cardsListDesc = `
List all profile cards.
`

func RegisterCreateCard(s *server.MCPServer, db *db.Engine) {
	s.AddTool(
		mcp.NewTool(
			"create_card",
			mcp.WithDescription(cardsCreateDesc),
			mcp.WithInputSchema[card.CardsCreateRequest](),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			cardsCreateReq, err := ConvertMCPRequest[card.CardsCreateRequest](req.GetArguments())
			if err != nil {
				return nil, fmt.Errorf("failed to convert arguments: %w", err)
			}
			fmt.Println("Received card creation request:", cardsCreateReq)
			rowCount, err := card.Create(ctx, db, cardsCreateReq)
			if err != nil {
				msg := fmt.Sprintf("failed to create card: %v", err)
				return mcp.NewToolResultError(msg), fmt.Errorf(msg)
			}
			if rowCount == 0 {
				return mcp.NewToolResultError("failed to create card"), fmt.Errorf("failed to create card")
			}
			return mcp.NewToolResultText(fmt.Sprintf("card created successfully. length: %d", rowCount)), nil
		},
	)
}
func RegisterListCards(s *server.MCPServer, db *db.Engine) {
	s.AddTool(
		mcp.NewTool(
			"list_cards",
			mcp.WithDescription(cardsListDesc),
			mcp.WithInputSchema[card.CardsListRequest](),
			mcp.WithOutputSchema[card.CardsListResponse](),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			cardsListReq, err := ConvertMCPRequest[card.CardsListRequest](req.GetArguments())
			if err != nil {
				return nil, fmt.Errorf("failed to convert arguments: %w", err)
			}
			fmt.Println("Received list cards request:", cardsListReq)
			response, err := card.List(ctx, db, cardsListReq)
			if err != nil {
				msg := fmt.Sprintf("failed to list cards: %v", err)
				return mcp.NewToolResultError(msg), fmt.Errorf(msg)
			}
			return mcp.NewToolResultJSON(response)
		},
	)
}
