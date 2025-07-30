package handlers

import "github.com/thunderjr/sptrans-mcp/internal/client"

// GlobalClient holds the SPTrans client instance for use by handlers
var GlobalClient *client.Client

// SetGlobalClient sets the global SPTrans client
func SetGlobalClient(c *client.Client) {
	GlobalClient = c
}