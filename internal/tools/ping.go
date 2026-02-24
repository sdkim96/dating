package tools

import (
	mcp "github.com/metoro-io/mcp-golang"
)

type PingArgs struct{}

func Register(s *mcp.Server) {
	s.RegisterTool("ping", "Health check", handlePing)
}

func handlePing(args PingArgs) (*mcp.ToolResponse, error) {
	return mcp.NewToolResponse(mcp.NewTextContent("pong")), nil
}
