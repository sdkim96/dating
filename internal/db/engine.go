package db

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Engine struct {
	Conn *sql.DB
}

type EngineOpt func(*Engine) error

func NewEngine(ctx context.Context, dsn string, opts ...EngineOpt) (*Engine, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	engine := &Engine{
		Conn: db,
	}

	for _, opt := range opts {
		if err := opt(engine); err != nil {
			return nil, err
		}
	}

	return engine, nil
}

func WithPing(ctx context.Context) EngineOpt {
	return func(e *Engine) error {
		if err := e.Conn.PingContext(ctx); err != nil {
			return err
		}
		return nil
	}
}

func WithMigrate(ctx context.Context) EngineOpt {
	return func(e *Engine) error {
		_, err := e.Conn.ExecContext(ctx, ddlSchema)
		return err
	}
}
