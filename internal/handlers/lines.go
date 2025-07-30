package handlers

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// SearchLinesParams defines the parameters for searching lines
type SearchLinesParams struct {
	SearchTerm string `json:"search_term" jsonschema:"The line name or number to search for (partial or complete)"`
}

// SearchLineByDirectionParams defines the parameters for searching lines by direction
type SearchLineByDirectionParams struct {
	SearchTerm string `json:"search_term" jsonschema:"The line code or identifier to search for"`
	Direction  int    `json:"direction" jsonschema:"The direction to search for (1 or 2)"`
}

// SearchLines handles the search_lines MCP tool
func SearchLines(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[SearchLinesParams]) (*mcp.CallToolResultFor[any], error) {
	if params.Arguments.SearchTerm == "" {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: "search_term parameter is required"}},
		}, nil
	}

	lines, err := GlobalClient.SearchLines(ctx, params.Arguments.SearchTerm)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to search lines: %v", err)}},
		}, nil
	}

	response := map[string]interface{}{
		"total_results": len(lines),
		"search_term":   params.Arguments.SearchTerm,
		"lines":         lines,
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("%+v", response)}},
		StructuredContent: response,
	}, nil
}

// SearchLineByDirection handles the search_line_by_direction MCP tool
func SearchLineByDirection(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[SearchLineByDirectionParams]) (*mcp.CallToolResultFor[any], error) {
	if params.Arguments.SearchTerm == "" {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: "search_term parameter is required"}},
		}, nil
	}

	if params.Arguments.Direction != 1 && params.Arguments.Direction != 2 {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: "direction parameter must be 1 or 2"}},
		}, nil
	}

	lines, err := GlobalClient.SearchLineByDirection(ctx, params.Arguments.SearchTerm, params.Arguments.Direction)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to search line by direction: %v", err)}},
		}, nil
	}

	response := map[string]interface{}{
		"total_results": len(lines),
		"search_term":   params.Arguments.SearchTerm,
		"direction":     params.Arguments.Direction,
		"lines":         lines,
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("%+v", response)}},
		StructuredContent: response,
	}, nil
}

