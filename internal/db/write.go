package db

import (
	"context"
)

type WriteEntry struct {
	Key      string
	Value    string
	Metadata string
}

func (e *Engine) Write(ctx context.Context, tenant, key, value string) error {
	_, err := e.Conn.ExecContext(ctx, writeQuery, tenant, key, value)

	return err
}

func (e *Engine) WriteBatch(ctx context.Context, tenant string, entries []WriteEntry) error {
	tx, err := e.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, writeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, e := range entries {
		if _, err := stmt.ExecContext(ctx, tenant, e.Key, e.Value, e.Metadata); err != nil {
			return err
		}
	}

	return tx.Commit()
}
