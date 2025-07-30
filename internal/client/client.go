package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/thunderjr/sptrans-mcp/internal/auth"
	"github.com/thunderjr/sptrans-mcp/internal/types"
)

const BaseURL = "https://api.olhovivo.sptrans.com.br/v2.1"

// Client wraps the SPTrans API with authentication
type Client struct {
	authManager *auth.Manager
	httpClient  *http.Client
}

// NewClient creates a new SPTrans API client
func NewClient(authManager *auth.Manager) *Client {
	return &Client{
		authManager: authManager,
		httpClient:  authManager.GetHTTPClient(),
	}
}

// makeRequest performs an authenticated HTTP request to the SPTrans API
func (c *Client) makeRequest(ctx context.Context, endpoint string, result interface{}) error {
	
	if err := c.authManager.EnsureAuthenticated(ctx); err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", BaseURL+endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "SPTrans-MCP-Server/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &types.APIError{
			Code:    resp.StatusCode,
			Message: "API request failed",
			Details: fmt.Sprintf("HTTP %d for endpoint %s", resp.StatusCode, endpoint),
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

// SearchLines searches for bus lines by name or number
func (c *Client) SearchLines(ctx context.Context, searchTerm string) ([]types.Line, error) {
	endpoint := fmt.Sprintf("/Linha/Buscar?termosBusca=%s", url.QueryEscape(searchTerm))
	var lines []types.Line
	if err := c.makeRequest(ctx, endpoint, &lines); err != nil {
		return nil, fmt.Errorf("failed to search lines: %w", err)
	}
	return lines, nil
}

// SearchLineByDirection searches for a specific line in a specific direction
func (c *Client) SearchLineByDirection(ctx context.Context, searchTerm string, direction int) ([]types.Line, error) {
	endpoint := fmt.Sprintf("/Linha/BuscarLinhaSentido?termosBusca=%s&sentido=%d", 
		url.QueryEscape(searchTerm), direction)
	var lines []types.Line
	if err := c.makeRequest(ctx, endpoint, &lines); err != nil {
		return nil, fmt.Errorf("failed to search line by direction: %w", err)
	}
	return lines, nil
}

// SearchStops searches for bus stops by name or address
func (c *Client) SearchStops(ctx context.Context, searchTerm string) ([]types.Stop, error) {
	endpoint := fmt.Sprintf("/Parada/Buscar?termosBusca=%s", url.QueryEscape(searchTerm))
	var stops []types.Stop
	if err := c.makeRequest(ctx, endpoint, &stops); err != nil {
		return nil, fmt.Errorf("failed to search stops: %w", err)
	}
	return stops, nil
}

// GetStopsByLine gets all stops served by a specific line
func (c *Client) GetStopsByLine(ctx context.Context, lineCode int) ([]types.Stop, error) {
	endpoint := fmt.Sprintf("/Parada/BuscarParadasPorLinha?codigoLinha=%d", lineCode)
	var stops []types.Stop
	if err := c.makeRequest(ctx, endpoint, &stops); err != nil {
		return nil, fmt.Errorf("failed to get stops by line: %w", err)
	}
	return stops, nil
}

// GetStopsByCorridor gets all stops in a specific corridor
func (c *Client) GetStopsByCorridor(ctx context.Context, corridorCode int) ([]types.Stop, error) {
	endpoint := fmt.Sprintf("/Parada/BuscarParadasPorCorredor?codigoCorredor=%d", corridorCode)
	var stops []types.Stop
	if err := c.makeRequest(ctx, endpoint, &stops); err != nil {
		return nil, fmt.Errorf("failed to get stops by corridor: %w", err)
	}
	return stops, nil
}

// GetCorridors retrieves all available corridors
func (c *Client) GetCorridors(ctx context.Context) ([]types.Corridor, error) {
	endpoint := "/Corredor"
	var corridors []types.Corridor
	if err := c.makeRequest(ctx, endpoint, &corridors); err != nil {
		return nil, fmt.Errorf("failed to get corridors: %w", err)
	}
	return corridors, nil
}

// GetCompanies retrieves all transport companies
func (c *Client) GetCompanies(ctx context.Context) (*types.Company, error) {
	endpoint := "/Empresa"
	var companies types.Company
	if err := c.makeRequest(ctx, endpoint, &companies); err != nil {
		return nil, fmt.Errorf("failed to get companies: %w", err)
	}
	return &companies, nil
}

// GetVehiclePositions gets real-time positions of all vehicles
func (c *Client) GetVehiclePositions(ctx context.Context) (*types.VehiclePositions, error) {
	endpoint := "/Posicao"
	var positions types.VehiclePositions
	if err := c.makeRequest(ctx, endpoint, &positions); err != nil {
		return nil, fmt.Errorf("failed to get vehicle positions: %w", err)
	}
	return &positions, nil
}

// GetVehiclePositionsByLine gets real-time positions of vehicles on a specific line
func (c *Client) GetVehiclePositionsByLine(ctx context.Context, lineCode int) (*types.VehiclePositions, error) {
	endpoint := fmt.Sprintf("/Posicao/Linha?codigoLinha=%d", lineCode)
	var positions types.VehiclePositions
	if err := c.makeRequest(ctx, endpoint, &positions); err != nil {
		return nil, fmt.Errorf("failed to get vehicle positions by line: %w", err)
	}
	return &positions, nil
}

// GetVehiclePositionsInGarage gets positions of vehicles currently in garage
func (c *Client) GetVehiclePositionsInGarage(ctx context.Context, companyCode, lineCode int) (*types.VehiclePositions, error) {
	endpoint := fmt.Sprintf("/Posicao/Garagem?codigoEmpresa=%d&codigoLinha=%d", companyCode, lineCode)
	var positions types.VehiclePositions
	if err := c.makeRequest(ctx, endpoint, &positions); err != nil {
		return nil, fmt.Errorf("failed to get vehicle positions in garage: %w", err)
	}
	return &positions, nil
}

// GetArrivalPredictions gets arrival predictions for vehicles at a specific stop and line
func (c *Client) GetArrivalPredictions(ctx context.Context, stopCode, lineCode int) (*types.ArrivalPrediction, error) {
	endpoint := fmt.Sprintf("/Previsao?codigoParada=%d&codigoLinha=%d", stopCode, lineCode)
	var predictions types.ArrivalPrediction
	if err := c.makeRequest(ctx, endpoint, &predictions); err != nil {
		return nil, fmt.Errorf("failed to get arrival predictions: %w", err)
	}
	return &predictions, nil
}

// GetArrivalPredictionsByLine gets all arrival predictions for a specific line
func (c *Client) GetArrivalPredictionsByLine(ctx context.Context, lineCode int) (*types.ArrivalPredictionsByLine, error) {
	endpoint := fmt.Sprintf("/Previsao/Linha?codigoLinha=%d", lineCode)
	var predictions types.ArrivalPredictionsByLine
	if err := c.makeRequest(ctx, endpoint, &predictions); err != nil {
		return nil, fmt.Errorf("failed to get arrival predictions by line: %w", err)
	}
	return &predictions, nil
}

// GetArrivalPredictionsByStop gets all arrival predictions for a specific stop
func (c *Client) GetArrivalPredictionsByStop(ctx context.Context, stopCode int) (*types.ArrivalPredictionsByLine, error) {
	endpoint := fmt.Sprintf("/Previsao/Parada?codigoParada=%d", stopCode)
	var predictions types.ArrivalPredictionsByLine
	if err := c.makeRequest(ctx, endpoint, &predictions); err != nil {
		return nil, fmt.Errorf("failed to get arrival predictions by stop: %w", err)
	}
	return &predictions, nil
}