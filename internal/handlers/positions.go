package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thunderjr/sptrans-mcp/internal/types"
)

// GetVehiclePositionsParams defines the parameters for getting all vehicle positions
type GetVehiclePositionsParams struct {
	// No parameters needed for getting all vehicle positions
}

// GetVehiclePositionsByLineParams defines the parameters for getting vehicle positions by line
type GetVehiclePositionsByLineParams struct {
	LineCode int `json:"line_code" jsonschema:"The line code to get vehicle positions for"`
}

// GetVehiclePositions handles the get_vehicle_positions MCP tool
func GetVehiclePositions(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[GetVehiclePositionsParams]) (*mcp.CallToolResultFor[any], error) {
	positions, err := GlobalClient.GetVehiclePositions(ctx)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to get vehicle positions: %v", err)}},
		}, nil
	}

	response := types.BuildGetVehiclePositionsResponse(*positions)

	responseJSON, err := json.Marshal(response)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to marshal response: %v", err)}},
		}, nil
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: string(responseJSON)}},
		StructuredContent: response,
	}, nil
}

// GetVehiclePositionsByLine handles the get_vehicle_positions_by_line MCP tool
func GetVehiclePositionsByLine(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[GetVehiclePositionsByLineParams]) (*mcp.CallToolResultFor[any], error) {
	if params.Arguments.LineCode <= 0 {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: "line_code parameter must be a positive integer"}},
		}, nil
	}

	positions, err := GlobalClient.GetVehiclePositionsByLine(ctx, params.Arguments.LineCode)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to get vehicle positions by line: %v", err)}},
		}, nil
	}

	response := types.BuildGetVehiclePositionsByLineResponse(params.Arguments.LineCode, *positions)

	responseJSON, err := json.Marshal(response)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to marshal response: %v", err)}},
		}, nil
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: string(responseJSON)}},
		StructuredContent: response,
	}, nil
}