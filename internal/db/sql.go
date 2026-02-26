package db

const ddlSchema = `
CREATE TABLE IF NOT EXISTS store (
	id SERIAL PRIMARY KEY,
	tenant VARCHAR(255) NOT NULL,
	key VARCHAR(255) NOT NULL,
	value TEXT NOT NULL,
	metadata JSONB NOT NULL DEFAULT '{}',
	created_at TIMESTAMPTZ DEFAULT NOW(),
	updated_at TIMESTAMPTZ DEFAULT NOW(),
	UNIQUE (tenant, key)
);
CREATE INDEX IF NOT EXISTS idx_store_tenant ON store (tenant);
CREATE INDEX IF NOT EXISTS idx_store_metadata ON store USING GIN (metadata);
`

const readAllQuery = `
SELECT key, value FROM store WHERE tenant = $1
`

const writeQuery = `
INSERT INTO store (tenant, key, value, metadata)
VALUES ($1, $2, $3, $4)
ON CONFLICT (tenant, key)
DO UPDATE SET value = $3, metadata = $4, updated_at = NOW()
`
