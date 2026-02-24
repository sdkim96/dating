package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func Register(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("ping",
			mcp.WithDescription("Health check"),
		),
		handlePing,
	)
}

func handlePing(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText("pong"), nil
}
