package types

import "time"

// Line represents a bus line in the SPTrans system
type Line struct {
	Code        int    `json:"cl"`        // Line code (unique identifier)
	IsCircular  bool   `json:"lc"`        // Is circular line
	Number      string `json:"lt"`        // Line number/name
	Direction   int    `json:"sl"`        // Direction (1 or 2)
	Type        int    `json:"tl"`        // Line type
	Origin      string `json:"tp"`        // Origin terminal
	Destination string `json:"ts"`        // Destination terminal
}

// Stop represents a bus stop in the SPTrans system
type Stop struct {
	Code      int     `json:"cp"`  // Stop code (unique identifier)
	Name      string  `json:"np"`  // Stop name
	Address   string  `json:"ed"`  // Stop address
	Latitude  float64 `json:"py"`  // Latitude
	Longitude float64 `json:"px"`  // Longitude
}

// Corridor represents a bus corridor in the SPTrans system
type Corridor struct {
	Code int    `json:"cc"` // Corridor code
	Name string `json:"nc"` // Corridor name
}

// Company represents a transport company
type Company struct {
	Hour string `json:"hr"` // Hour of data
	Data []struct {
		Area      int `json:"a"` // Area code
		Companies []struct {
			Area int    `json:"a"` // Area code
			Code int    `json:"c"` // Company code
			Name string `json:"n"` // Company name
		} `json:"e"`
	} `json:"e"`
}

// Vehicle represents a vehicle position
type Vehicle struct {
	ID          int       `json:"p"`   // Vehicle identifier
	Accessible  bool      `json:"a"`   // Is accessible vehicle
	LastUpdate  time.Time `json:"ta"`  // Last update timestamp
	Latitude    float64   `json:"py"`  // Latitude
	Longitude   float64   `json:"px"`  // Longitude
}

// VehiclePositions represents the response from vehicle position endpoints
type VehiclePositions struct {
	Hour  string `json:"hr"` // Data timestamp
	Lines []struct {
		Identifier  string    `json:"c"`   // Line identifier
		Code        int       `json:"cl"`  // Line code
		Direction   int       `json:"sl"`  // Direction
		Origin      string    `json:"lt0"` // Origin terminal
		Destination string    `json:"lt1"` // Destination terminal
		VehicleQty  int       `json:"qv"`  // Number of vehicles
		Vehicles    []Vehicle `json:"vs"`  // Vehicles data
	} `json:"l"`
}

// ArrivalPrediction represents arrival prediction data
type ArrivalPrediction struct {
	Hour string `json:"hr"` // Current time
	Stop struct {
		Code      int     `json:"cp"` // Stop code
		Name      string  `json:"np"` // Stop name
		Latitude  float64 `json:"py"` // Stop latitude
		Longitude float64 `json:"px"` // Stop longitude
		Lines     []struct {
			Identifier    string `json:"c"`   // Line identifier
			Code          int    `json:"cl"`  // Line code
			Direction     int    `json:"sl"`  // Direction
			Origin        string `json:"lt0"` // Origin terminal
			Destination   string `json:"lt1"` // Destination terminal
			VehicleQty    int    `json:"qv"`  // Number of vehicles
			Predictions   []struct {
				VehicleID   string    `json:"p"`  // Vehicle identifier
				ArrivalTime string    `json:"t"`  // Predicted arrival time
				Accessible  bool      `json:"a"`  // Is accessible vehicle
				LastUpdate  time.Time `json:"ta"` // Last position update
				Latitude    float64   `json:"py"` // Current vehicle latitude
				Longitude   float64   `json:"px"` // Current vehicle longitude
			} `json:"vs"`
		} `json:"l"`
	} `json:"p"`
}

// ArrivalPredictionsByLine represents predictions for all stops on a line
type ArrivalPredictionsByLine struct {
	Hour  string `json:"hr"` // Current time
	Stops []struct {
		Code      int     `json:"cp"` // Stop code
		Name      string  `json:"np"` // Stop name
		Latitude  float64 `json:"py"` // Stop latitude
		Longitude float64 `json:"px"` // Stop longitude
		Lines     []struct {
			Identifier    string `json:"c"`   // Line identifier
			Code          int    `json:"cl"`  // Line code
			Direction     int    `json:"sl"`  // Direction
			Origin        string `json:"lt0"` // Origin terminal
			Destination   string `json:"lt1"` // Destination terminal
			VehicleQty    int    `json:"qv"`  // Number of vehicles
			Predictions   []struct {
				VehicleID   string    `json:"p"`  // Vehicle identifier
				ArrivalTime string    `json:"t"`  // Predicted arrival time
				Accessible  bool      `json:"a"`  // Is accessible vehicle
				LastUpdate  time.Time `json:"ta"` // Last position update
				Latitude    float64   `json:"py"` // Current vehicle latitude
				Longitude   float64   `json:"px"` // Current vehicle longitude
			} `json:"vs"`
		} `json:"l"`
	} `json:"ps"`
}

// APIError represents an error response from the SPTrans API
type APIError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e APIError) Error() string {
	if e.Details != "" {
		return e.Message + ": " + e.Details
	}
	return e.Message
}