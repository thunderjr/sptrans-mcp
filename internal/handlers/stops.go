package handlers

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// SearchStopsParams defines the parameters for searching stops
type SearchStopsParams struct {
	SearchTerm string `json:"search_term" jsonschema:"The stop name or address to search for (partial or complete)"`
}

// GetStopsByLineParams defines the parameters for getting stops by line
type GetStopsByLineParams struct {
	LineCode int `json:"line_code" jsonschema:"The line code to get stops for"`
}

// GetStopsByCorridorParams defines the parameters for getting stops by corridor
type GetStopsByCorridorParams struct {
	CorridorCode int `json:"corridor_code" jsonschema:"The corridor code to get stops for"`
}

// SearchStops handles the search_stops MCP tool
func SearchStops(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[SearchStopsParams]) (*mcp.CallToolResultFor[any], error) {
	if params.Arguments.SearchTerm == "" {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: "search_term parameter is required"}},
		}, nil
	}

	stops, err := GlobalClient.SearchStops(ctx, params.Arguments.SearchTerm)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to search stops: %v", err)}},
		}, nil
	}

	response := map[string]interface{}{
		"total_results": len(stops),
		"search_term":   params.Arguments.SearchTerm,
		"stops":         stops,
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("%+v", response)}},
		StructuredContent: response,
	}, nil
}

// GetStopsByLine handles the get_stops_by_line MCP tool
func GetStopsByLine(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[GetStopsByLineParams]) (*mcp.CallToolResultFor[any], error) {
	if params.Arguments.LineCode <= 0 {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: "line_code parameter must be a positive integer"}},
		}, nil
	}

	stops, err := GlobalClient.GetStopsByLine(ctx, params.Arguments.LineCode)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to get stops by line: %v", err)}},
		}, nil
	}

	response := map[string]interface{}{
		"total_results": len(stops),
		"line_code":     params.Arguments.LineCode,
		"stops":         stops,
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("%+v", response)}},
		StructuredContent: response,
	}, nil
}

// GetStopsByCorridor handles the get_stops_by_corridor MCP tool
func GetStopsByCorridor(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[GetStopsByCorridorParams]) (*mcp.CallToolResultFor[any], error) {
	if params.Arguments.CorridorCode <= 0 {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: "corridor_code parameter must be a positive integer"}},
		}, nil
	}

	stops, err := GlobalClient.GetStopsByCorridor(ctx, params.Arguments.CorridorCode)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to get stops by corridor: %v", err)}},
		}, nil
	}

	response := map[string]interface{}{
		"total_results":  len(stops),
		"corridor_code":  params.Arguments.CorridorCode,
		"stops":          stops,
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("%+v", response)}},
		StructuredContent: response,
	}, nil
}