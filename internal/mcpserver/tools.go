package mcpserver

import (
	"context"
	"fmt"

	"github.com/Zolet-hash/mcp-golang/internal/currency"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type GreetingInput struct {
	Name string `json:"name" jsonschema:"Name of the person to greet"`
}

type GreetingOutput struct {
	Message string `json:"message"`
}

type ConvertInput struct {
	Amount float64 `json:"amount" jsonschema:"Amount to convert"`
	From   string  `json:"from" jsonschema:"Source currency code"`
	To     string  `json:"to" jsonschema:"Target currency code"`
}

type ConvertOutput struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Result float64 `json:"result"`
}

func registerTools(
	server *mcp.Server,
	converter *currency.Converter,
) {
	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "greeting",
			Description: "Greet a person by name.",
		},
		func(
			ctx context.Context,
			req *mcp.CallToolRequest,
			input GreetingInput,
		) (*mcp.CallToolResult, GreetingOutput, error) {
			message := fmt.Sprintf("Hello, %s!", input.Name)

			return nil, GreetingOutput{
				Message: message,
			}, nil
		},
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name: "convert_currency",
			Description: "Convert currencies using illustrative " +
				"fixed exchange rates. These rates are not live.",
		},
		func(
			ctx context.Context,
			req *mcp.CallToolRequest,
			input ConvertInput,
		) (*mcp.CallToolResult, ConvertOutput, error) {
			result, err := converter.Convert(
				input.Amount,
				input.From,
				input.To,
			)
			if err != nil {
				return nil, ConvertOutput{}, err
			}

			return nil, ConvertOutput{
				Amount: input.Amount,
				From:   input.From,
				To:     input.To,
				Result: result,
			}, nil
		},
	)
}
