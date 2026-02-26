package app

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/server"
	"github.com/sdkim96/dating/internal/config"
	"github.com/sdkim96/dating/internal/db"
	"github.com/sdkim96/dating/internal/tools"
)

type App struct {
	Server *server.MCPServer
	DB     *db.Engine
}

func NewApp(cfg *config.Config) *App {

	fmt.Println("[INIT] Initializing database connection...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbEngine, err := db.NewEngine(
		ctx,
		cfg.DB.DSN(),
		db.WithPing(ctx),
		db.WithMigrate(ctx),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("[INIT] Initializing MCP server and registering tools...")
	s := server.NewMCPServer("dating", "0.1.0")

	app := &App{Server: s, DB: dbEngine}
	tools.RegisterPing(s)
	tools.RegisterCreateCard(s, dbEngine)
	tools.RegisterListCards(s, dbEngine)

	return app

}

func (app *App) Close() error {
	if app.DB == nil {
		return nil
	}
	return app.DB.Conn.Close()
}
