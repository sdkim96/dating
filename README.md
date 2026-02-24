# dating

MCP server built with Go and [mcp-go](https://github.com/mark3labs/mcp-go).

## Project Structure

```
cmd/dating/main.go          # Entry point
internal/server/server.go   # Server setup (stdio & http)
internal/tools/ping.go      # Tool handlers
```

## Run

```sh
# stdio (default)
go run ./cmd/dating

# http stateless on :8080 (no session required)
go run ./cmd/dating -mode http

# http stateful (requires initialize + Mcp-Session-Id)
go run ./cmd/dating -mode http-stateful

# custom port
go run ./cmd/dating -mode http -addr :3000
```

## Docker

```sh
docker build -t dating .
docker run -p 8080:8080 dating
```

## HTTP Endpoint

The server listens for JSON-RPC POST requests at `/api`.

### Stateless (default http)

```sh
curl -X POST http://localhost:8080/api \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "ping",
      "arguments": {}
    }
  }'
```

### Stateful

Requires an `initialize` call first:

```sh
# 1. Initialize and get Mcp-Session-Id from response header
curl -v -X POST http://localhost:8080/api \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "initialize",
    "params": {
      "protocolVersion": "2025-03-26",
      "capabilities": {},
      "clientInfo": { "name": "curl", "version": "1.0" }
    }
  }'

# 2. Call tool with session ID
curl -X POST http://localhost:8080/api \
  -H "Content-Type: application/json" \
  -H "Mcp-Session-Id: <session-id-from-step-1>" \
  -d '{
    "jsonrpc": "2.0",
    "id": 2,
    "method": "tools/call",
    "params": {
      "name": "ping",
      "arguments": {}
    }
  }'
```

## MCP Server Card

```json
{
  "mcpServers": {
    "dating": {
      "type": "streamable-http",
      "url": "http://localhost:8080/api"
    }
  }
}
```
