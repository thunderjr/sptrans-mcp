# SPTrans Olho Vivo API
---

## API Documentation

- **Base URL**: `https://api.olhovivo.sptrans.com.br/v2.1`

### Authentication

#### Authenticate Token
Validates your access token and establishes a session.

**Endpoint**: `POST /Login/Autenticar`

**Parameters**:
- `token` (string, required): Your API access token

**Request Example**:
```http
POST /Login/Autenticar?token=your_access_token_here
```

**Response**:
```json
true
```

**Response Codes**:
- `true`: Authentication successful
- `false`: Authentication failed

---

### Line Operations

#### Search Lines
Search for bus lines by name or number (partial or complete).

**Endpoint**: `GET /Linha/Buscar`

**Parameters**:
- `termosBusca` (string, required): Line name or number (partial or complete)

**Request Example**:
```http
GET /Linha/Buscar?termosBusca=8000
```

**Response Example**:
```json
[
  {
    "cl": 1273,
    "lc": false,
    "lt": "8000",
    "sl": 1,
    "tl": 10,
    "tp": "PCA.RAMOS DE AZEVEDO",
    "ts": "TERMINAL LAPA"
  }
]
```

**Response Fields**:
- `cl` (int): Line code (unique identifier)
- `lc` (bool): Is circular line
- `lt` (string): Line number/name
- `sl` (int): Direction (1 or 2)
- `tl` (int): Line type
- `tp` (string): Origin terminal
- `ts` (string): Destination terminal

#### Search Line by Direction
Search for a specific line in a specific direction.

**Endpoint**: `GET /Linha/BuscarLinhaSentido`

**Parameters**:
- `termosBusca` (string, required): Line code
- `sentido` (byte, required): Direction (1 or 2)

**Request Example**:
```http
GET /Linha/BuscarLinhaSentido?termosBusca=8000&sentido=1
```

---

### Stop Operations

#### Search Stops
Search for bus stops by name or address (partial or complete).

**Endpoint**: `GET /Parada/Buscar`

**Parameters**:
- `termosBusca` (string, required): Stop name or address (partial or complete)

**Request Example**:
```http
GET /Parada/Buscar?termosBusca=AFONSO BRAZ
```

**Response Example**:
```json
[
  {
    "cp": 340015329,
    "np": "AFONSO BRAZ B/C1",
    "ed": "R ARMINDA/ R BALTHAZAR DA VEIGA",
    "py": -23.592938,
    "px": -46.672727
  }
]
```

**Response Fields**:
- `cp` (int): Stop code (unique identifier)
- `np` (string): Stop name
- `ed` (string): Stop address
- `py` (double): Latitude
- `px` (double): Longitude

#### Get Stops by Line
Get all stops served by a specific line.

**Endpoint**: `GET /Parada/BuscarParadasPorLinha`

**Parameters**:
- `codigoLinha` (int, required): Line code

**Request Example**:
```http
GET /Parada/BuscarParadasPorLinha?codigoLinha=1273
```

#### Get Stops by Corridor
Get all stops in a specific corridor.

**Endpoint**: `GET /Parada/BuscarParadasPorCorredor`

**Parameters**:
- `codigoCorredor` (int, required): Corridor code

**Request Example**:
```http
GET /Parada/BuscarParadasPorCorredor?codigoCorredor=8
```

---

### Corridor Operations

#### Get All Corridors
Retrieve all available corridors.

**Endpoint**: `GET /Corredor`

**Response Example**:
```json
[
  {
    "cc": 8,
    "nc": "Campo Limpo"
  }
]
```

**Response Fields**:
- `cc` (int): Corridor code
- `nc` (string): Corridor name

---

### Company Operations

#### Get All Companies
Retrieve all transport companies.

**Endpoint**: `GET /Empresa`

**Response Example**:
```json
[
  {
    "hr": "11:20",
    "e": [
      {
        "a": 1,
        "e": [
          {
            "a": 1,
            "c": 999,
            "n": "COMPANY NAME"
          }
        ]
      }
    ]
  }
]
```

**Response Fields**:
- `hr` (string): Hour of data
- `e` (array): Companies data
- `a` (int): Area code
- `c` (int): Company code
- `n` (string): Company name

---

### Vehicle Position Operations

#### Get All Vehicle Positions
Get real-time positions of all vehicles.

**Endpoint**: `GET /Posicao`

**Response Example**:
```json
{
  "hr": "11:30",
  "l": [
    {
      "c": "5015-10",
      "cl": 33887,
      "sl": 2,
      "lt0": "METRÔ JABAQUARA",
      "lt1": "JD. SÃO JORGE",
      "qv": 1,
      "vs": [
        {
          "p": 68021,
          "a": true,
          "ta": "2017-05-12T14:30:37Z",
          "py": -23.678712500000003,
          "px": -46.65674
        }
      ]
    }
  ]
}
```

**Response Fields**:
- `hr` (string): Data timestamp
- `l` (array): Lines data
- `c` (string): Line identifier
- `cl` (int): Line code
- `sl` (int): Direction
- `lt0` (string): Origin terminal
- `lt1` (string): Destination terminal
- `qv` (int): Number of vehicles
- `vs` (array): Vehicles data
- `p` (int): Vehicle identifier
- `a` (bool): Is accessible vehicle
- `ta` (string): Last update timestamp
- `py` (double): Latitude
- `px` (double): Longitude

#### Get Vehicle Positions by Line
Get real-time positions of vehicles on a specific line.

**Endpoint**: `GET /Posicao/Linha`

**Parameters**:
- `codigoLinha` (int, required): Line code

**Request Example**:
```http
GET /Posicao/Linha?codigoLinha=1273
```

#### Get Vehicle Positions in Garage
Get positions of vehicles currently in garage.

**Endpoint**: `GET /Posicao/Garagem`

**Parameters**:
- `codigoEmpresa` (int, optional): Company code (0 for all)
- `codigoLinha` (int, optional): Line code (0 for all)

**Request Example**:
```http
GET /Posicao/Garagem?codigoEmpresa=0&codigoLinha=0
```

---

### Arrival Prediction Operations

#### Get Arrival Predictions
Get arrival predictions for vehicles at a specific stop and line.

**Endpoint**: `GET /Previsao`

**Parameters**:
- `codigoParada` (int, required): Stop code
- `codigoLinha` (int, required): Line code

**Request Example**:
```http
GET /Previsao?codigoParada=4200953&codigoLinha=1989
```

**Response Example**:
```json
{
  "hr": "20:09",
  "p": {
    "cp": 4200953,
    "np": "PARADA ROBERTO SELMI DEI B/C",
    "py": -23.675901,
    "px": -46.752812,
    "l": [
      {
        "c": "7021-10",
        "cl": 1989,
        "sl": 1,
        "lt0": "TERM. JOÃO DIAS",
        "lt1": "JD. MARACÁ",
        "qv": 1,
        "vs": [
          {
            "p": "74558",
            "t": "23:11",
            "a": true,
            "ta": "2017-05-07T23:09:05Z",
            "py": -23.67603,
            "px": -46.75891166666667
          }
        ]
      }
    ]
  }
}
```

**Response Fields**:
- `hr` (string): Current time
- `p` (object): Stop information
- `cp` (int): Stop code
- `np` (string): Stop name
- `py` (double): Stop latitude
- `px` (double): Stop longitude
- `l` (array): Lines serving this stop
- `c` (string): Line identifier
- `cl` (int): Line code
- `sl` (int): Direction
- `lt0` (string): Origin terminal
- `lt1` (string): Destination terminal
- `qv` (int): Number of vehicles
- `vs` (array): Vehicle predictions
- `p` (string): Vehicle identifier
- `t` (string): Predicted arrival time
- `a` (bool): Is accessible vehicle
- `ta` (string): Last position update
- `py` (double): Current vehicle latitude
- `px` (double): Current vehicle longitude

#### Get Arrival Predictions by Line
Get all arrival predictions for a specific line.

**Endpoint**: `GET /Previsao/Linha`

**Parameters**:
- `codigoLinha` (int, required): Line code

**Request Example**:
```http
GET /Previsao/Linha?codigoLinha=1273
```

#### Get Arrival Predictions by Stop
Get all arrival predictions for a specific stop.

**Endpoint**: `GET /Previsao/Parada`

**Parameters**:
- `codigoParada` (int, required): Stop code

**Request Example**:
```http
GET /Previsao/Parada?codigoParada=4200953
```

---

### KMZ Operations

#### Get All Routes KMZ
Download KMZ file with all bus routes.

**Endpoint**: `GET /KMZ`

**Optional Parameters**:
- `sentido` (string, optional): Direction filter

#### Get Express Routes KMZ
Download KMZ file with express bus routes only.

**Endpoint**: `GET /KMZ/BC`

**Optional Parameters**:
- `sentido` (string, optional): Direction filter

#### Get Corridor Routes KMZ
Download KMZ file with corridor routes.

**Endpoint**: `GET /KMZ/Corredor`

**Optional Parameters**:
- `sentido` (string, optional): Direction filter

#### Get Corridor Express Routes KMZ
Download KMZ file with corridor express routes.

**Endpoint**: `GET /KMZ/Corredor/BC`

**Optional Parameters**:
- `sentido` (string, optional): Direction filter

#### Get Other Routes KMZ
Download KMZ file with other routes.

**Endpoint**: `GET /KMZ/OutrasVias`

**Optional Parameters**:
- `sentido` (string, optional): Direction filter

#### Get Other Express Routes KMZ
Download KMZ file with other express routes.

**Endpoint**: `GET /KMZ/OutrasVias/BC`

**Optional Parameters**:
- `sentido` (string, optional): Direction filter

---

## Implementation Examples

### JavaScript/Node.js Example

```javascript
class SPTransAPI {
  constructor(token) {
    this.token = token;
    this.baseURL = 'https://api.olhovivo.sptrans.com.br/v2.1';
    this.authenticated = false;
  }

  async authenticate() {
    try {
      const response = await fetch(`${this.baseURL}/Login/Autenticar?token=${this.token}`, {
        method: 'POST'
      });

      const result = await response.json();
      this.authenticated = result === true;
      return this.authenticated;
    } catch (error) {
      console.error('Authentication failed:', error);
      return false;
    }
  }

  async searchLines(searchTerm) {
    if (!this.authenticated) {
      throw new Error('Not authenticated. Call authenticate() first.');
    }

    try {
      const response = await fetch(`${this.baseURL}/Linha/Buscar?termosBusca=${encodeURIComponent(searchTerm)}`);
      return await response.json();
    } catch (error) {
      console.error('Error searching lines:', error);
      throw error;
    }
  }

  async getVehiclePositions(lineCode = null) {
    if (!this.authenticated) {
      throw new Error('Not authenticated. Call authenticate() first.');
    }

    try {
      const endpoint = lineCode
        ? `${this.baseURL}/Posicao/Linha?codigoLinha=${lineCode}`
        : `${this.baseURL}/Posicao`;

      const response = await fetch(endpoint);
      return await response.json();
    } catch (error) {
      console.error('Error getting vehicle positions:', error);
      throw error;
    }
  }

  async getArrivalPredictions(stopCode, lineCode) {
    if (!this.authenticated) {
      throw new Error('Not authenticated. Call authenticate() first.');
    }

    try {
      const response = await fetch(`${this.baseURL}/Previsao?codigoParada=${stopCode}&codigoLinha=${lineCode}`);
      return await response.json();
    } catch (error) {
      console.error('Error getting arrival predictions:', error);
      throw error;
    }
  }
}

// Usage example
async function example() {
  const api = new SPTransAPI('your_token_here');

  // Authenticate
  const authenticated = await api.authenticate();
  if (!authenticated) {
    console.error('Failed to authenticate');
    return;
  }

  // Search for lines
  const lines = await api.searchLines('8000');
  console.log('Found lines:', lines);

  // Get vehicle positions for a specific line
  if (lines.length > 0) {
    const positions = await api.getVehiclePositions(lines[0].cl);
    console.log('Vehicle positions:', positions);
  }
}
```

### Python Example

```python
import requests
import json
from typing import Optional, List, Dict, Any

class SPTransAPI:
    def __init__(self, token: str):
        self.token = token
        self.base_url = 'https://api.olhovivo.sptrans.com.br/v2.1'
        self.session = requests.Session()
        self.authenticated = False

    def authenticate(self) -> bool:
        """Authenticate with the API"""
        try:
            response = self.session.post(f'{self.base_url}/Login/Autenticar?token={self.token}')
            response.raise_for_status()

            result = response.json()
            self.authenticated = result is True
            return self.authenticated
        except requests.RequestException as e:
            print(f'Authentication failed: {e}')
            return False

    def search_lines(self, search_term: str) -> List[Dict[str, Any]]:
        """Search for bus lines"""
        if not self.authenticated:
            raise Exception('Not authenticated. Call authenticate() first.')

        try:
            response = self.session.get(f'{self.base_url}/Linha/Buscar?termosBusca={search_term}')
            response.raise_for_status()
            return response.json()
        except requests.RequestException as e:
            print(f'Error searching lines: {e}')
            raise

    def get_vehicle_positions(self, line_code: Optional[int] = None) -> Dict[str, Any]:
        """Get vehicle positions"""
        if not self.authenticated:
            raise Exception('Not authenticated. Call authenticate() first.')

        try:
            endpoint = f'{self.base_url}/Posicao/Linha?codigoLinha={line_code}' if line_code else f'{self.base_url}/Posicao'
            response = self.session.get(endpoint)
            response.raise_for_status()
            return response.json()
        except requests.RequestException as e:
            print(f'Error getting vehicle positions: {e}')
            raise

    def get_arrival_predictions(self, stop_code: int, line_code: int) -> Dict[str, Any]:
        """Get arrival predictions"""
        if not self.authenticated:
            raise Exception('Not authenticated. Call authenticate() first.')

        try:
            response = self.session.get(f'{self.base_url}/Previsao?codigoParada={stop_code}&codigoLinha={line_code}')
            response.raise_for_status()
            return response.json()
        except requests.RequestException as e:
            print(f'Error getting arrival predictions: {e}')
            raise

# Usage example
def main():
    api = SPTransAPI('your_token_here')

    # Authenticate
    if not api.authenticate():
        print('Failed to authenticate')
        return

    # Search for lines
    lines = api.search_lines('8000')
    print(f'Found {len(lines)} lines')

    # Get vehicle positions for first line
    if lines:
        positions = api.get_vehicle_positions(lines[0]['cl'])
        print(f'Found {len(positions.get("vs", []))} vehicles')

if __name__ == '__main__':
    main()
```

---

## Best Practices

### 1. Authentication Management
- Always validate authentication before making API calls
- Implement token refresh mechanisms if needed
- Store tokens securely (environment variables, secure storage)
- Handle authentication failures gracefully

### 2. Error Handling
- Implement comprehensive error handling for all API calls
- Handle network timeouts and connection issues
- Log errors appropriately for debugging
- Provide meaningful error messages to users

### 3. Rate Limiting & Performance
- Respect API rate limits to avoid blocking
- Implement exponential backoff for failed requests
- Cache frequently accessed data when appropriate
- Use connection pooling for better performance

### 4. Data Processing
- Validate API responses before processing
- Handle missing or null fields gracefully
- Implement data transformation layers
- Consider data normalization for consistency

### 5. Monitoring & Logging
- Log all API interactions for debugging
- Monitor API usage and performance metrics
- Set up alerts for API failures or unusual patterns
- Track data freshness and accuracy

### 6. Security Considerations
- Never expose API tokens in client-side code
- Use HTTPS for all API communications
- Implement proper input sanitization
- Follow security best practices for data handling

### 7. Testing Strategy
- Create comprehensive unit tests for API interactions
- Implement integration tests with real API endpoints
- Test error scenarios and edge cases
- Use mock data for development and testing

### 8. Documentation & Maintenance
- Keep integration documentation up to date
- Document all custom configurations and modifications
- Maintain version compatibility tracking
- Create operational runbooks for common issues
