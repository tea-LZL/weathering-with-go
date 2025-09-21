# Weathering with Go 🌤️

A modern, high-performance weather API built with Go and the Gin web framework. Get current weather conditions and forecasts for any location worldwide using the OpenWeatherMap API.

## ✨ Features

- **Current Weather**: Get real-time weather data for any location
- **Weather Forecasts**: 5-day weather forecasts with 3-hour intervals
- **Multiple Units**: Support for metric, imperial, and Kelvin units
- **RESTful API**: Clean, well-documented REST endpoints
- **Error Handling**: Comprehensive error handling with detailed responses
- **Rate Limiting**: Built-in protection against API abuse
- **CORS Support**: Cross-origin resource sharing enabled
- **Health Checks**: Monitor API health and status
- **Middleware**: Security headers, logging, and request tracking

## 🚀 Quick Start

### Prerequisites

- Go 1.19 or higher
- OpenWeatherMap API key (free at [openweathermap.org](https://openweathermap.org/api))

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/tea-LZL/weathering-with-go.git
   cd weathering-with-go
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set environment variables**
   ```bash
   export OPENWEATHERMAP_API_KEY="your_api_key_here"
   export PORT="8080"  # Optional, defaults to 8080
   ```

4. **Run the server**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`

## 📖 API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Authentication
No authentication required. The OpenWeatherMap API key is configured server-side.

### Endpoints

#### GET /health
Health check endpoint to verify the API is running.

**Response:**
```json
{
  "status": "healthy",
  "service": "weathering-with-go"
}
```

#### GET /weather/current
Get current weather for a location.

**Parameters:**
- `location` (required): City name, state code, and country code (e.g., "London,UK" or "New York,NY,US")
- `units` (optional): Temperature units - `metric` (default), `imperial`, or `kelvin`

**Example:**
```bash
curl "http://localhost:8080/api/v1/weather/current?location=London,UK&units=metric"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "location": {
      "name": "London",
      "country": "GB",
      "latitude": 51.5074,
      "longitude": -0.1278
    },
    "current": {
      "temperature": 15.5,
      "feels_like": 14.8,
      "humidity": 72,
      "pressure": 1013.2,
      "wind_speed": 3.6,
      "wind_direction": 230,
      "condition": "Clouds",
      "description": "Scattered Clouds",
      "icon": "03d",
      "cloud_cover": 40,
      "last_updated": "2025-09-21T10:30:00Z"
    },
    "request_time": "2025-09-21T10:30:15Z"
  }
}
```

#### POST /weather/current
Get current weather using JSON request body.

**Request Body:**
```json
{
  "location": "London,UK",
  "units": "metric"
}
```

**Response:** Same as GET endpoint

#### GET /weather/forecast
Get weather forecast for a location.

**Parameters:**
- `location` (required): City name, state code, and country code
- `units` (optional): Temperature units - `metric` (default), `imperial`, or `kelvin`
- `days` (optional): Number of forecast days (1-5, default: 5)

**Example:**
```bash
curl "http://localhost:8080/api/v1/weather/forecast?location=Tokyo,JP&units=metric&days=3"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "location": {
      "name": "Tokyo",
      "country": "JP",
      "latitude": 35.6762,
      "longitude": 139.6503
    },
    "forecast": [
      {
        "date": "2025-09-22T00:00:00Z",
        "max_temperature": 24.5,
        "min_temperature": 18.2,
        "avg_temperature": 21.3,
        "condition": "Clear",
        "description": "Clear Sky",
        "icon": "01d",
        "humidity": 65,
        "wind_speed": 2.1,
        "precipitation": 0,
        "chance_of_rain": 0
      }
    ],
    "request_time": "2025-09-21T10:30:15Z"
  }
}
```

#### POST /weather/forecast
Get weather forecast using JSON request body.

**Request Body:**
```json
{
  "location": "Tokyo,JP",
  "units": "metric",
  "days": 3
}
```

**Response:** Same as GET endpoint

### Error Responses

All errors follow a consistent format:

```json
{
  "success": false,
  "error": {
    "error": "Bad Request",
    "code": 400,
    "message": "Location is required"
  }
}
```

**Common Error Codes:**
- `400`: Bad Request - Invalid parameters
- `401`: Unauthorized - Invalid API key
- `404`: Not Found - Location not found
- `429`: Too Many Requests - Rate limit exceeded
- `500`: Internal Server Error - Server error
- `503`: Service Unavailable - External API unavailable

## ⚙️ Configuration

### Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `OPENWEATHERMAP_API_KEY` | Yes | - | Your OpenWeatherMap API key |
| `PORT` | No | `8080` | Server port |
| `HOST` | No | `0.0.0.0` | Server host |
| `ENVIRONMENT` | No | `development` | Environment (development/production) |
| `LOG_LEVEL` | No | `info` | Log level (debug/info/warn/error) |

### Example .env file
```env
OPENWEATHERMAP_API_KEY=your_api_key_here
PORT=8080
HOST=0.0.0.0
ENVIRONMENT=development
LOG_LEVEL=info
```

## 🏗️ Project Structure

```
weathering-with-go/
├── config/
│   └── config.go           # Configuration management
├── handlers/
│   ├── routes.go          # Route definitions
│   └── weather.go         # Weather request handlers
├── middleware/
│   └── middleware.go      # HTTP middleware
├── models/
│   ├── openweather.go     # OpenWeatherMap API models
│   └── weather.go         # Internal data models
├── services/
│   └── weather.go         # Weather service logic
├── utils/
│   └── errors.go          # Error handling utilities
├── go.mod                 # Go module dependencies
├── go.sum                 # Dependency checksums
├── main.go               # Application entry point
└── README.md             # This file
```

## 🔧 Development

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
go build -o weathering-with-go main.go
```

### Docker Support
```bash
# Build image
docker build -t weathering-with-go .

# Run container
docker run -p 8080:8080 -e OPENWEATHERMAP_API_KEY=your_key weathering-with-go
```

## 📊 API Limits

- **OpenWeatherMap Free Tier**: 1,000 calls/day, 60 calls/minute
- **Forecast**: Up to 5 days (OpenWeatherMap limitation)
- **Rate Limiting**: Built-in protection to prevent API abuse

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [OpenWeatherMap](https://openweathermap.org/) for providing the weather data API
- [Gin Framework](https://gin-gonic.com/) for the excellent HTTP web framework
- [Go Community](https://golang.org/) for the amazing programming language

## 📞 Support

If you have any questions or need help, please:
1. Check the [issues](https://github.com/tea-LZL/weathering-with-go/issues) for existing solutions
2. Create a new issue with detailed information
3. Contact the maintainers

---

**Happy weather forecasting!** 🌈
