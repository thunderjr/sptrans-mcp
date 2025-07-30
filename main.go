package main

import (
	"context"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thunderjr/sptrans-mcp/internal/auth"
	"github.com/thunderjr/sptrans-mcp/internal/client"
	"github.com/thunderjr/sptrans-mcp/internal/handlers"
)

func main() {
	ctx := context.Background()

	// Get the SPTrans API token from environment
	token := os.Getenv("SPTRANS_PAT")
	if token == "" {
		log.Fatal("SPTRANS_PAT environment variable is required")
	}

	// Create authentication manager
	authManager := auth.NewManager(token)

	// Authenticate on startup
	if err := authManager.Authenticate(ctx); err != nil {
		log.Fatalf("Failed to authenticate with SPTrans API: %v", err)
	}
	log.Println("Successfully authenticated with SPTrans API")

	// Create SPTrans client
	sptransClient := client.NewClient(authManager)

	// Set the global client for handlers to use
	handlers.SetGlobalClient(sptransClient)

	// Create MCP server
	server := mcp.NewServer(&mcp.Implementation{Name: "sptrans-mcp", Version: "1.0.0"}, nil)

	// Register line operation tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "search_lines",
		Description: "Search for bus lines by name or number (partial or complete)",
	}, handlers.SearchLines)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "search_line_by_direction",
		Description: "Search for a specific line in a specific direction",
	}, handlers.SearchLineByDirection)

	// Register stop operation tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "search_stops",
		Description: "Search for bus stops by name or address (partial or complete)",
	}, handlers.SearchStops)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_stops_by_line",
		Description: "Get all stops served by a specific line",
	}, handlers.GetStopsByLine)

	// Register vehicle position tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_vehicle_positions",
		Description: "Get real-time positions of all vehicles",
	}, handlers.GetVehiclePositions)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_vehicle_positions_by_line",
		Description: "Get real-time positions of vehicles on a specific line",
	}, handlers.GetVehiclePositionsByLine)

	// Register arrival prediction tools (core for forecasting)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_arrival_predictions",
		Description: "Get arrival predictions for vehicles at a specific stop and line",
	}, handlers.GetArrivalPredictions)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_arrival_predictions_by_line",
		Description: "Get all arrival predictions for a specific line",
	}, handlers.GetArrivalPredictionsByLine)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_arrival_predictions_by_stop",
		Description: "Get all arrival predictions for a specific stop",
	}, handlers.GetArrivalPredictionsByStop)

	log.Println("SPTrans MCP Server starting...")
	log.Println("Available tools:")
	log.Println("  - search_lines: Search for bus lines")
	log.Println("  - search_line_by_direction: Search line by direction")
	log.Println("  - search_stops: Search for bus stops")
	log.Println("  - get_stops_by_line: Get stops for a line")
	log.Println("  - get_vehicle_positions: Get all vehicle positions")
	log.Println("  - get_vehicle_positions_by_line: Get vehicle positions by line")
	log.Println("  - get_arrival_predictions: Get arrival predictions")
	log.Println("  - get_arrival_predictions_by_line: Get predictions by line")
	log.Println("  - get_arrival_predictions_by_stop: Get predictions by stop")

	// Run the server over stdin/stdout
	if err := server.Run(ctx, mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}
}
