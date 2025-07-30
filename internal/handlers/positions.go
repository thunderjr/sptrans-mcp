package handlers

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
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

	totalVehicles := 0
	for _, line := range positions.Lines {
		totalVehicles += line.VehicleQty
	}

	response := map[string]interface{}{
		"timestamp":     positions.Hour,
		"total_vehicles": totalVehicles,
		"total_lines":    len(positions.Lines),
		"positions":      positions,
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("%+v", response)}},
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

	totalVehicles := 0
	for _, line := range positions.Lines {
		totalVehicles += line.VehicleQty
	}

	response := map[string]interface{}{
		"timestamp":     positions.Hour,
		"line_code":     params.Arguments.LineCode,
		"total_vehicles": totalVehicles,
		"total_lines":    len(positions.Lines),
		"positions":      positions,
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("%+v", response)}},
		StructuredContent: response,
	}, nil
}