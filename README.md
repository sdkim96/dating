# dating

MCP server built with Go and [mcp-golang](https://github.com/metoro-io/mcp-golang).

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

# http on :8080
go run ./cmd/dating -mode http

# http on custom port
go run ./cmd/dating -mode http -addr :3000
```

## HTTP Endpoint

When running in `http` mode, the server listens for JSON-RPC POST requests at `/mcp`.

## Adding Tools

Create a new file in `internal/tools/`, define a handler, and register it in `Register()`:

```go
type MyArgs struct {
    Name string `json:"name" jsonschema:"description=User name,required"`
}

func handleMyTool(args MyArgs) (*mcp.ToolResponse, error) {
    return mcp.NewToolResponse(mcp.NewTextContent("hello " + args.Name)), nil
}
```

```go
func Register(s *mcp.Server) {
    s.RegisterTool("ping", "Health check", handlePing)
    s.RegisterTool("my_tool", "My tool description", handleMyTool)
}
```
