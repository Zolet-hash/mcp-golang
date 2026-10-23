package mcpserver

import (
	"github.com/Zolet-hash/mcp-golang/internal/currency"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func New(converter *currency.Converter) *mcp.Server {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name: "go-currency-server",
			Version: "1.0.0",
		},
		nil,
	)
	registerTools(server, converter)
	return server
}
