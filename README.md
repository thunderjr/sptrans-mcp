# SPTrans MCP Server

MCP server for São Paulo SPTrans Olho Vivo API - bus lines, stops, positions, and arrival predictions.

## Setup

1. Get your SPTrans API token from [SPTrans Developer Portal](https://www.sptrans.com.br/desenvolvedores/)

2. Add to Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "sptrans": {
      "command": "go",
      "args": ["run", "/path/to/sptrans-mcp/main.go"],
      "env": {
        "SPTRANS_PAT": "your_api_token_here"
      }
    }
  }
}
```

## Tools

- `search_lines` - Find bus lines by name/number
- `search_stops` - Find bus stops by name/address  
- `get_stops_by_line` - Get stops for a specific line
- `get_vehicle_positions` - Get real-time vehicle positions
- `get_arrival_predictions` - Get bus arrival predictions