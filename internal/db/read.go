package db

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

type MetadataFilter struct {
	Field string
	Value string
}

type ReadAllOption struct {
	OrderBy  string
	Limit    int
	Offset   int
	Metadata []MetadataFilter
}

var allowedOrderBy = map[string]string{
	"key":        "key",
	"created_at": "created_at",
	"updated_at": "updated_at",
}

// ReadAll retrieves all key-value pairs to a given tenant.
// This key-value pairs is marshaled as JSON string.
// The returned JSON string is:
// {"{tenantName}": {"{key1}": "{value1}", "{key2}": "{value2}", ...}}
// If no entries exist for the tenant, it returns an empty string without error.
func (e *Engine) ReadAll(
	ctx context.Context,
	tenant string,
	opt *ReadAllOption,
) (string, error) {
	query := readAllQuery
	args := []any{tenant}
	idx := 2

	if opt != nil {
		for _, m := range opt.Metadata {
			query += " AND metadata ->> $" + strconv.Itoa(idx) + " ILIKE $" + strconv.Itoa(idx+1)
			args = append(args, m.Field, "%"+m.Value+"%")
			idx += 2
		}
		if opt.OrderBy != "" {
			if col, ok := allowedOrderBy[opt.OrderBy]; ok {
				query += " ORDER BY " + col
			} else {
				return "", fmt.Errorf("invalid order by column: %s", opt.OrderBy)
			}
		}
		if opt.Limit > 0 {
			query += " LIMIT $" + strconv.Itoa(idx)
			args = append(args, opt.Limit)
			idx++
		}
		if opt.Offset > 0 {
			query += " OFFSET $" + strconv.Itoa(idx)
			args = append(args, opt.Offset)
			idx++
		}
	}

	rows, err := e.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	kvs := make(map[string]string)
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return "", err
		}
		kvs[key] = value
	}
	if err := rows.Err(); err != nil {
		return "", err
	}
	if len(kvs) == 0 {
		return "", nil
	}

	result := map[string]map[string]string{tenant: kvs}
	b, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
