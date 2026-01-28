package types

// Conversion functions to transform SPTrans structs to clean JSON response structs

// ConvertLine converts a Line struct to LineResponse
func ConvertLine(line Line) LineResponse {
	return LineResponse{
		Code:        line.Code,
		IsCircular:  line.IsCircular,
		Number:      line.Number,
		Direction:   line.Direction,
		Type:        line.Type,
		Origin:      line.Origin,
		Destination: line.Destination,
	}
}

// ConvertLines converts a slice of Line structs to LineResponse structs
func ConvertLines(lines []Line) []LineResponse {
	result := make([]LineResponse, len(lines))
	for i, line := range lines {
		result[i] = ConvertLine(line)
	}
	return result
}

// ConvertStop converts a Stop struct to StopResponse
func ConvertStop(stop Stop) StopResponse {
	return StopResponse{
		Code:      stop.Code,
		Name:      stop.Name,
		Address:   stop.Address,
		Latitude:  stop.Latitude,
		Longitude: stop.Longitude,
	}
}

// ConvertStops converts a slice of Stop structs to StopResponse structs
func ConvertStops(stops []Stop) []StopResponse {
	result := make([]StopResponse, len(stops))
	for i, stop := range stops {
		result[i] = ConvertStop(stop)
	}
	return result
}

// ConvertCorridor converts a Corridor struct to CorridorResponse
func ConvertCorridor(corridor Corridor) CorridorResponse {
	return CorridorResponse{
		Code: corridor.Code,
		Name: corridor.Name,
	}
}

// ConvertCorridors converts a slice of Corridor structs to CorridorResponse structs
func ConvertCorridors(corridors []Corridor) []CorridorResponse {
	result := make([]CorridorResponse, len(corridors))
	for i, corridor := range corridors {
		result[i] = ConvertCorridor(corridor)
	}
	return result
}

// ConvertVehicle converts a Vehicle struct to VehicleResponse
func ConvertVehicle(vehicle Vehicle) VehicleResponse {
	return VehicleResponse{
		ID:          vehicle.ID,
		Accessible:  vehicle.Accessible,
		LastUpdate:  vehicle.LastUpdate,
		Latitude:    vehicle.Latitude,
		Longitude:   vehicle.Longitude,
	}
}

// ConvertVehicles converts a slice of Vehicle structs to VehicleResponse structs
func ConvertVehicles(vehicles []Vehicle) []VehicleResponse {
	result := make([]VehicleResponse, len(vehicles))
	for i, vehicle := range vehicles {
		result[i] = ConvertVehicle(vehicle)
	}
	return result
}

// ConvertVehiclePositions converts VehiclePositions struct to VehiclePositionsResponse
func ConvertVehiclePositions(positions VehiclePositions) VehiclePositionsResponse {
	lines := make([]LineWithVehiclesResponse, len(positions.Lines))
	for i, line := range positions.Lines {
		lines[i] = LineWithVehiclesResponse{
			Identifier:   line.Identifier,
			Code:         line.Code,
			Direction:    line.Direction,
			Origin:       line.Origin,
			Destination:  line.Destination,
			VehicleCount: line.VehicleQty,
			Vehicles:     ConvertVehicles(line.Vehicles),
		}
	}

	return VehiclePositionsResponse{
		Timestamp: positions.Hour,
		Lines:     lines,
	}
}

// ConvertPrediction converts a prediction struct to PredictionResponse
func ConvertPrediction(pred struct {
	VehicleID   string    `json:"p"`
	ArrivalTime string    `json:"t"`
	Accessible  bool      `json:"a"`
	LastUpdate  interface{} `json:"ta"`
	Latitude    float64   `json:"py"`
	Longitude   float64   `json:"px"`
}) PredictionResponse {
	return PredictionResponse{
		VehicleID:   pred.VehicleID,
		ArrivalTime: pred.ArrivalTime,
		Accessible:  pred.Accessible,
		// Handle the LastUpdate field conversion if needed
		Latitude:    pred.Latitude,
		Longitude:   pred.Longitude,
	}
}

// ConvertArrivalPrediction converts ArrivalPrediction struct to ArrivalPredictionResponse
func ConvertArrivalPrediction(predictions ArrivalPrediction) ArrivalPredictionResponse {
	lines := make([]LineWithPredictionsResponse, len(predictions.Stop.Lines))
	for i, line := range predictions.Stop.Lines {
		preds := make([]PredictionResponse, len(line.Predictions))
		for j, pred := range line.Predictions {
			preds[j] = PredictionResponse{
				VehicleID:   pred.VehicleID,
				ArrivalTime: pred.ArrivalTime,
				Accessible:  pred.Accessible,
				LastUpdate:  pred.LastUpdate,
				Latitude:    pred.Latitude,
				Longitude:   pred.Longitude,
			}
		}

		lines[i] = LineWithPredictionsResponse{
			Identifier:    line.Identifier,
			Code:          line.Code,
			Direction:     line.Direction,
			Origin:        line.Origin,
			Destination:   line.Destination,
			VehicleCount:  line.VehicleQty,
			Predictions:   preds,
		}
	}

	stop := StopWithPredictionsResponse{
		Code:      predictions.Stop.Code,
		Name:      predictions.Stop.Name,
		Latitude:  predictions.Stop.Latitude,
		Longitude: predictions.Stop.Longitude,
		Lines:     lines,
	}

	return ArrivalPredictionResponse{
		Timestamp: predictions.Hour,
		Stop:      stop,
	}
}

// ConvertArrivalPredictionsByLine converts ArrivalPredictionsByLine struct to ArrivalPredictionsByLineResponse
func ConvertArrivalPredictionsByLine(predictions ArrivalPredictionsByLine) ArrivalPredictionsByLineResponse {
	stops := make([]StopWithPredictionsResponse, len(predictions.Stops))
	for i, stop := range predictions.Stops {
		lines := make([]LineWithPredictionsResponse, len(stop.Lines))
		for j, line := range stop.Lines {
			preds := make([]PredictionResponse, len(line.Predictions))
			for k, pred := range line.Predictions {
				preds[k] = PredictionResponse{
					VehicleID:   pred.VehicleID,
					ArrivalTime: pred.ArrivalTime,
					Accessible:  pred.Accessible,
					LastUpdate:  pred.LastUpdate,
					Latitude:    pred.Latitude,
					Longitude:   pred.Longitude,
				}
			}

			lines[j] = LineWithPredictionsResponse{
				Identifier:    line.Identifier,
				Code:          line.Code,
				Direction:     line.Direction,
				Origin:        line.Origin,
				Destination:   line.Destination,
				VehicleCount:  line.VehicleQty,
				Predictions:   preds,
			}
		}

		stops[i] = StopWithPredictionsResponse{
			Code:      stop.Code,
			Name:      stop.Name,
			Latitude:  stop.Latitude,
			Longitude: stop.Longitude,
			Lines:     lines,
		}
	}

	return ArrivalPredictionsByLineResponse{
		Timestamp: predictions.Hour,
		Stops:     stops,
	}
}

// Response builder functions

// BuildSearchLinesResponse builds a SearchLinesResponse
func BuildSearchLinesResponse(totalResults int, searchTerm string, lines []Line) SearchLinesResponse {
	return SearchLinesResponse{
		TotalResults: totalResults,
		SearchTerm:   searchTerm,
		Lines:        ConvertLines(lines),
	}
}

// BuildSearchStopsResponse builds a SearchStopsResponse
func BuildSearchStopsResponse(totalResults int, searchTerm string, stops []Stop) SearchStopsResponse {
	return SearchStopsResponse{
		TotalResults: totalResults,
		SearchTerm:   searchTerm,
		Stops:        ConvertStops(stops),
	}
}

// BuildGetStopsByLineResponse builds a GetStopsByLineResponse
func BuildGetStopsByLineResponse(totalResults int, lineCode int, stops []Stop) GetStopsByLineResponse {
	return GetStopsByLineResponse{
		TotalResults: totalResults,
		LineCode:     lineCode,
		Stops:        ConvertStops(stops),
	}
}

// BuildGetStopsByCorridorResponse builds a GetStopsByCorridorResponse
func BuildGetStopsByCorridorResponse(totalResults int, corridorCode int, stops []Stop) GetStopsByCorridorResponse {
	return GetStopsByCorridorResponse{
		TotalResults:  totalResults,
		CorridorCode:  corridorCode,
		Stops:         ConvertStops(stops),
	}
}

// BuildGetVehiclePositionsResponse builds a GetVehiclePositionsResponse
func BuildGetVehiclePositionsResponse(positions VehiclePositions) GetVehiclePositionsResponse {
	totalVehicles := 0
	for _, line := range positions.Lines {
		totalVehicles += line.VehicleQty
	}

	convertedPositions := ConvertVehiclePositions(positions)

	return GetVehiclePositionsResponse{
		Timestamp:     positions.Hour,
		TotalVehicles: totalVehicles,
		TotalLines:    len(positions.Lines),
		Positions:     convertedPositions,
	}
}

// BuildGetVehiclePositionsByLineResponse builds a GetVehiclePositionsByLineResponse
func BuildGetVehiclePositionsByLineResponse(lineCode int, positions VehiclePositions) GetVehiclePositionsByLineResponse {
	totalVehicles := 0
	for _, line := range positions.Lines {
		totalVehicles += line.VehicleQty
	}

	convertedPositions := ConvertVehiclePositions(positions)

	return GetVehiclePositionsByLineResponse{
		Timestamp:     positions.Hour,
		LineCode:      lineCode,
		TotalVehicles: totalVehicles,
		TotalLines:    len(positions.Lines),
		Positions:     convertedPositions,
	}
}

// BuildGetArrivalPredictionsResponse builds a GetArrivalPredictionsResponse
func BuildGetArrivalPredictionsResponse(stopCode, lineCode int, predictions ArrivalPrediction) GetArrivalPredictionsResponse {
	totalPredictions := 0
	for _, line := range predictions.Stop.Lines {
		totalPredictions += len(line.Predictions)
	}

	convertedPredictions := ConvertArrivalPrediction(predictions)

	return GetArrivalPredictionsResponse{
		Timestamp:        predictions.Hour,
		StopCode:         stopCode,
		LineCode:         lineCode,
		TotalPredictions: totalPredictions,
		Predictions:      convertedPredictions,
	}
}

// BuildGetArrivalPredictionsByLineResponse builds a GetArrivalPredictionsByLineResponse
func BuildGetArrivalPredictionsByLineResponse(lineCode int, predictions ArrivalPredictionsByLine) GetArrivalPredictionsByLineResponse {
	totalPredictions := 0
	for _, stop := range predictions.Stops {
		for _, line := range stop.Lines {
			totalPredictions += len(line.Predictions)
		}
	}

	convertedPredictions := ConvertArrivalPredictionsByLine(predictions)

	return GetArrivalPredictionsByLineResponse{
		Timestamp:        predictions.Hour,
		LineCode:         lineCode,
		TotalPredictions: totalPredictions,
		TotalStops:       len(predictions.Stops),
		Predictions:      convertedPredictions,
	}
}

// BuildGetArrivalPredictionsByStopResponse builds a GetArrivalPredictionsByStopResponse
func BuildGetArrivalPredictionsByStopResponse(stopCode int, predictions ArrivalPredictionsByLine) GetArrivalPredictionsByStopResponse {
	totalPredictions := 0
	for _, stop := range predictions.Stops {
		for _, line := range stop.Lines {
			totalPredictions += len(line.Predictions)
		}
	}

	convertedPredictions := ConvertArrivalPredictionsByLine(predictions)

	return GetArrivalPredictionsByStopResponse{
		Timestamp:        predictions.Hour,
		StopCode:         stopCode,
		TotalPredictions: totalPredictions,
		TotalStops:       len(predictions.Stops),
		Predictions:      convertedPredictions,
	}
}