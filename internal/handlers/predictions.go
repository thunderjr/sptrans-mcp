package handlers

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetArrivalPredictionsParams defines the parameters for getting arrival predictions
type GetArrivalPredictionsParams struct {
	StopCode int `json:"stop_code" jsonschema:"The stop code to get predictions for"`
	LineCode int `json:"line_code" jsonschema:"The line code to get predictions for"`
}

// GetArrivalPredictionsByLineParams defines the parameters for getting predictions by line
type GetArrivalPredictionsByLineParams struct {
	LineCode int `json:"line_code" jsonschema:"The line code to get all predictions for"`
}

// GetArrivalPredictionsByStopParams defines the parameters for getting predictions by stop
type GetArrivalPredictionsByStopParams struct {
	StopCode int `json:"stop_code" jsonschema:"The stop code to get all predictions for"`
}

// GetArrivalPredictions handles the get_arrival_predictions MCP tool
func GetArrivalPredictions(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[GetArrivalPredictionsParams]) (*mcp.CallToolResultFor[any], error) {
	if params.Arguments.StopCode <= 0 {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: "stop_code parameter must be a positive integer"}},
		}, nil
	}

	if params.Arguments.LineCode <= 0 {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: "line_code parameter must be a positive integer"}},
		}, nil
	}

	predictions, err := GlobalClient.GetArrivalPredictions(ctx, params.Arguments.StopCode, params.Arguments.LineCode)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to get arrival predictions: %v", err)}},
		}, nil
	}

	totalPredictions := 0
	for _, line := range predictions.Stop.Lines {
		totalPredictions += len(line.Predictions)
	}

	response := map[string]interface{}{
		"timestamp":        predictions.Hour,
		"stop_code":        params.Arguments.StopCode,
		"line_code":        params.Arguments.LineCode,
		"total_predictions": totalPredictions,
		"predictions":       predictions,
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("%+v", response)}},
		StructuredContent: response,
	}, nil
}

// GetArrivalPredictionsByLine handles the get_arrival_predictions_by_line MCP tool
func GetArrivalPredictionsByLine(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[GetArrivalPredictionsByLineParams]) (*mcp.CallToolResultFor[any], error) {
	if params.Arguments.LineCode <= 0 {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: "line_code parameter must be a positive integer"}},
		}, nil
	}

	predictions, err := GlobalClient.GetArrivalPredictionsByLine(ctx, params.Arguments.LineCode)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to get arrival predictions by line: %v", err)}},
		}, nil
	}

	totalPredictions := 0
	for _, stop := range predictions.Stops {
		for _, line := range stop.Lines {
			totalPredictions += len(line.Predictions)
		}
	}

	response := map[string]interface{}{
		"timestamp":        predictions.Hour,
		"line_code":        params.Arguments.LineCode,
		"total_predictions": totalPredictions,
		"total_stops":       len(predictions.Stops),
		"predictions":       predictions,
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("%+v", response)}},
		StructuredContent: response,
	}, nil
}

// GetArrivalPredictionsByStop handles the get_arrival_predictions_by_stop MCP tool
func GetArrivalPredictionsByStop(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[GetArrivalPredictionsByStopParams]) (*mcp.CallToolResultFor[any], error) {
	if params.Arguments.StopCode <= 0 {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: "stop_code parameter must be a positive integer"}},
		}, nil
	}

	predictions, err := GlobalClient.GetArrivalPredictionsByStop(ctx, params.Arguments.StopCode)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Failed to get arrival predictions by stop: %v", err)}},
		}, nil
	}

	totalPredictions := 0
	for _, stop := range predictions.Stops {
		for _, line := range stop.Lines {
			totalPredictions += len(line.Predictions)
		}
	}

	response := map[string]interface{}{
		"timestamp":        predictions.Hour,
		"stop_code":        params.Arguments.StopCode,
		"total_predictions": totalPredictions,
		"total_stops":       len(predictions.Stops),
		"predictions":       predictions,
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("%+v", response)}},
		StructuredContent: response,
	}, nil
}