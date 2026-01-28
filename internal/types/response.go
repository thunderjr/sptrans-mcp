package types

import "time"

// Clean JSON response structs with readable field names

// LineResponse represents a bus line with clean JSON field names
type LineResponse struct {
	Code        int    `json:"code"`        // Line code (unique identifier)
	IsCircular  bool   `json:"is_circular"` // Is circular line
	Number      string `json:"number"`      // Line number/name
	Direction   int    `json:"direction"`   // Direction (1 or 2)
	Type        int    `json:"type"`        // Line type
	Origin      string `json:"origin"`      // Origin terminal
	Destination string `json:"destination"` // Destination terminal
}

// StopResponse represents a bus stop with clean JSON field names
type StopResponse struct {
	Code      int     `json:"code"`      // Stop code (unique identifier)
	Name      string  `json:"name"`      // Stop name
	Address   string  `json:"address"`   // Stop address
	Latitude  float64 `json:"latitude"`  // Latitude
	Longitude float64 `json:"longitude"` // Longitude
}

// CorridorResponse represents a bus corridor with clean JSON field names
type CorridorResponse struct {
	Code int    `json:"code"` // Corridor code
	Name string `json:"name"` // Corridor name
}

// VehicleResponse represents a vehicle position with clean JSON field names
type VehicleResponse struct {
	ID          int       `json:"id"`           // Vehicle identifier
	Accessible  bool      `json:"accessible"`   // Is accessible vehicle
	LastUpdate  time.Time `json:"last_update"`  // Last update timestamp
	Latitude    float64   `json:"latitude"`     // Latitude
	Longitude   float64   `json:"longitude"`    // Longitude
}

// LineWithVehiclesResponse represents a line with its vehicles
type LineWithVehiclesResponse struct {
	Identifier   string            `json:"identifier"`    // Line identifier
	Code         int               `json:"code"`          // Line code
	Direction    int               `json:"direction"`     // Direction
	Origin       string            `json:"origin"`        // Origin terminal
	Destination  string            `json:"destination"`   // Destination terminal
	VehicleCount int               `json:"vehicle_count"` // Number of vehicles
	Vehicles     []VehicleResponse `json:"vehicles"`      // Vehicles data
}

// VehiclePositionsResponse represents vehicle positions with clean JSON field names
type VehiclePositionsResponse struct {
	Timestamp string                     `json:"timestamp"` // Data timestamp
	Lines     []LineWithVehiclesResponse `json:"lines"`     // Lines with vehicles
}

// PredictionResponse represents arrival prediction data with clean JSON field names
type PredictionResponse struct {
	VehicleID   string    `json:"vehicle_id"`   // Vehicle identifier
	ArrivalTime string    `json:"arrival_time"` // Predicted arrival time
	Accessible  bool      `json:"accessible"`   // Is accessible vehicle
	LastUpdate  time.Time `json:"last_update"`  // Last position update
	Latitude    float64   `json:"latitude"`     // Current vehicle latitude
	Longitude   float64   `json:"longitude"`    // Current vehicle longitude
}

// LineWithPredictionsResponse represents a line with its predictions
type LineWithPredictionsResponse struct {
	Identifier    string               `json:"identifier"`      // Line identifier
	Code          int                  `json:"code"`            // Line code
	Direction     int                  `json:"direction"`       // Direction
	Origin        string               `json:"origin"`          // Origin terminal
	Destination   string               `json:"destination"`     // Destination terminal
	VehicleCount  int                  `json:"vehicle_count"`   // Number of vehicles
	Predictions   []PredictionResponse `json:"predictions"`     // Predictions data
}

// StopWithPredictionsResponse represents a stop with its predictions
type StopWithPredictionsResponse struct {
	Code      int                           `json:"code"`       // Stop code
	Name      string                        `json:"name"`       // Stop name
	Latitude  float64                       `json:"latitude"`   // Stop latitude
	Longitude float64                       `json:"longitude"`  // Stop longitude
	Lines     []LineWithPredictionsResponse `json:"lines"`      // Lines with predictions
}

// ArrivalPredictionResponse represents arrival prediction data with clean JSON field names
type ArrivalPredictionResponse struct {
	Timestamp string                      `json:"timestamp"` // Current time
	Stop      StopWithPredictionsResponse `json:"stop"`      // Stop with predictions
}

// ArrivalPredictionsByLineResponse represents predictions for all stops on a line
type ArrivalPredictionsByLineResponse struct {
	Timestamp string                        `json:"timestamp"` // Current time
	Stops     []StopWithPredictionsResponse `json:"stops"`     // Stops with predictions
}

// SearchLinesResponse represents the response for line searches
type SearchLinesResponse struct {
	TotalResults int            `json:"total_results"` // Number of results found
	SearchTerm   string         `json:"search_term"`   // Original search term
	Lines        []LineResponse `json:"lines"`         // Found lines
}

// SearchStopsResponse represents the response for stop searches
type SearchStopsResponse struct {
	TotalResults int            `json:"total_results"` // Number of results found
	SearchTerm   string         `json:"search_term"`   // Original search term
	Stops        []StopResponse `json:"stops"`         // Found stops
}

// GetStopsByLineResponse represents the response for getting stops by line
type GetStopsByLineResponse struct {
	TotalResults int            `json:"total_results"` // Number of results found
	LineCode     int            `json:"line_code"`     // Line code used
	Stops        []StopResponse `json:"stops"`         // Found stops
}

// GetStopsByCorridorResponse represents the response for getting stops by corridor
type GetStopsByCorridorResponse struct {
	TotalResults  int            `json:"total_results"`  // Number of results found
	CorridorCode  int            `json:"corridor_code"`  // Corridor code used
	Stops         []StopResponse `json:"stops"`          // Found stops
}

// GetVehiclePositionsResponse represents the response for vehicle positions
type GetVehiclePositionsResponse struct {
	Timestamp     string                       `json:"timestamp"`      // Data timestamp
	TotalVehicles int                          `json:"total_vehicles"` // Total number of vehicles
	TotalLines    int                          `json:"total_lines"`    // Total number of lines
	Positions     VehiclePositionsResponse     `json:"positions"`      // Vehicle positions data
}

// GetVehiclePositionsByLineResponse represents the response for vehicle positions by line
type GetVehiclePositionsByLineResponse struct {
	Timestamp     string                       `json:"timestamp"`      // Data timestamp
	LineCode      int                          `json:"line_code"`      // Line code used
	TotalVehicles int                          `json:"total_vehicles"` // Total number of vehicles
	TotalLines    int                          `json:"total_lines"`    // Total number of lines
	Positions     VehiclePositionsResponse     `json:"positions"`      // Vehicle positions data
}

// GetArrivalPredictionsResponse represents the response for arrival predictions
type GetArrivalPredictionsResponse struct {
	Timestamp        string                    `json:"timestamp"`         // Data timestamp
	StopCode         int                       `json:"stop_code"`         // Stop code used
	LineCode         int                       `json:"line_code"`         // Line code used
	TotalPredictions int                       `json:"total_predictions"` // Total number of predictions
	Predictions      ArrivalPredictionResponse `json:"predictions"`       // Predictions data
}

// GetArrivalPredictionsByLineResponse represents the response for predictions by line
type GetArrivalPredictionsByLineResponse struct {
	Timestamp        string                           `json:"timestamp"`         // Data timestamp
	LineCode         int                              `json:"line_code"`         // Line code used
	TotalPredictions int                              `json:"total_predictions"` // Total number of predictions
	TotalStops       int                              `json:"total_stops"`       // Total number of stops
	Predictions      ArrivalPredictionsByLineResponse `json:"predictions"`       // Predictions data
}

// GetArrivalPredictionsByStopResponse represents the response for predictions by stop
type GetArrivalPredictionsByStopResponse struct {
	Timestamp        string                           `json:"timestamp"`         // Data timestamp
	StopCode         int                              `json:"stop_code"`         // Stop code used
	TotalPredictions int                              `json:"total_predictions"` // Total number of predictions
	TotalStops       int                              `json:"total_stops"`       // Total number of stops
	Predictions      ArrivalPredictionsByLineResponse `json:"predictions"`       // Predictions data
}