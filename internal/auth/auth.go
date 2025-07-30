package auth

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"

	"github.com/thunderjr/sptrans-mcp/internal/types"
)

const (
	AuthEndpoint = "https://api.olhovivo.sptrans.com.br/v2.1/Login/Autenticar"
	TokenTimeout = 30 * time.Minute // SPTrans tokens typically expire after 30 minutes
)

// Manager handles SPTrans API authentication
type Manager struct {
	token         string
	client        *http.Client
	authenticated bool
	lastAuth      time.Time
	mu            sync.RWMutex
}

// NewManager creates a new authentication manager
func NewManager(token string) *Manager {
	// Create cookie jar to maintain session cookies after authentication
	jar, err := cookiejar.New(nil)
	if err != nil {
		// Fallback to client without cookie jar if creation fails
		jar = nil
	}
	
	// Create HTTP client with cookie jar to maintain session
	client := &http.Client{
		Timeout: 30 * time.Second,
		Jar:     jar,
	}
	
	return &Manager{
		token:  token,
		client: client,
	}
}

// SetHTTPClient allows setting a custom HTTP client
func (m *Manager) SetHTTPClient(client *http.Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.client = client
}

// IsAuthenticated checks if the current session is authenticated and not expired
func (m *Manager) IsAuthenticated() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	return m.authenticated && time.Since(m.lastAuth) < TokenTimeout
}

// Authenticate performs authentication with the SPTrans API
func (m *Manager) Authenticate(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if already authenticated and not expired
	if m.authenticated && time.Since(m.lastAuth) < TokenTimeout {
		return nil
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s?token=%s", AuthEndpoint, m.token), nil)
	if err != nil {
		return fmt.Errorf("failed to create auth request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "SPTrans-MCP-Server/1.0")

	resp, err := m.client.Do(req)
	if err != nil {
		return fmt.Errorf("authentication request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &types.APIError{
			Code:    resp.StatusCode,
			Message: "Authentication failed",
			Details: fmt.Sprintf("HTTP %d", resp.StatusCode),
		}
	}

	// Read the response body to properly determine the result
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	bodyStr := string(body)
	// SPTrans API returns plain "true" or "false" as response body
	result := bodyStr == "true"

	if !result {
		m.authenticated = false
		return &types.APIError{
			Code:    401,
			Message: "Invalid authentication token",
			Details: "SPTrans API returned false for authentication",
		}
	}

	m.authenticated = true
	m.lastAuth = time.Now()
	
	return nil
}

// EnsureAuthenticated ensures the session is authenticated, re-authenticating if necessary
func (m *Manager) EnsureAuthenticated(ctx context.Context) error {
	if m.IsAuthenticated() {
		return nil
	}
	
	return m.Authenticate(ctx)
}

// GetHTTPClient returns an HTTP client that can be used for authenticated requests
func (m *Manager) GetHTTPClient() *http.Client {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.client
}