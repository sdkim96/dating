package server

import (
	"github.com/mark3labs/mcp-go/server"
	"github.com/sdkim96/dating/internal/tools"
)

func newMCPServer() *server.MCPServer {
	s := server.NewMCPServer("dating", "0.1.0")
	tools.Register(s)
	return s
}

func RunStdio() error {
	return server.ServeStdio(newMCPServer())
}

func RunHTTPStateless(addr string) error {
	s := newMCPServer()
	httpServer := server.NewStreamableHTTPServer(s,
		server.WithEndpointPath("/api"),
		server.WithStateLess(true),
	)
	return httpServer.Start(addr)
}

func RunHTTPStateful(addr string) error {
	s := newMCPServer()
	httpServer := server.NewStreamableHTTPServer(s,
		server.WithEndpointPath("/api"),
		server.WithStateful(true),
	)
	return httpServer.Start(addr)
}
