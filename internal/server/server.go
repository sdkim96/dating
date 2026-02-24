package server

import (
	mcp "github.com/metoro-io/mcp-golang"
	mcphttp "github.com/metoro-io/mcp-golang/transport/http"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"github.com/sdkim96/dating/internal/tools"
)

func NewStdio() *mcp.Server {
	t := stdio.NewStdioServerTransport()
	s := mcp.NewServer(t, mcp.WithName("dating"), mcp.WithVersion("0.1.0"))
	tools.Register(s)
	return s
}

func NewHTTP(addr string) *mcp.Server {
	t := mcphttp.NewHTTPTransport("/mcp").WithAddr(addr)
	s := mcp.NewServer(
		t,
		mcp.WithName("dating"),
		mcp.WithVersion("0.1.0"),
	)
	tools.Register(s)
	return s
}
