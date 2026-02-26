package app

import (
	"github.com/mark3labs/mcp-go/server"
	"github.com/sdkim96/dating/internal/config"
)

func RunStdio(cfg *config.Config) error {
	return server.ServeStdio(NewApp(cfg).Server)
}

func RunHTTPStateless(cfg *config.Config, addr string) error {
	s := NewApp(cfg).Server
	httpServer := server.NewStreamableHTTPServer(s,
		server.WithEndpointPath("/api"),
		server.WithStateLess(true),
	)
	return httpServer.Start(addr)
}

func RunHTTPStateful(cfg *config.Config, addr string) error {
	s := NewApp(cfg).Server
	httpServer := server.NewStreamableHTTPServer(s,
		server.WithEndpointPath("/api"),
		server.WithStateful(true),
	)
	return httpServer.Start(addr)
}
