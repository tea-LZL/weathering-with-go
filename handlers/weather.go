package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"weathering-with-go/models"
	"weathering-with-go/services"
	"weathering-with-go/utils"
)

// WeatherHandler handles weather-related HTTP requests
type WeatherHandler struct {
	weatherService *services.WeatherService
}

// NewWeatherHandler creates a new weather handler instance
func NewWeatherHandler(weatherService *services.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		weatherService: weatherService,
	}
}

// GetCurrentWeather handles GET /weather/current requests
func (h *WeatherHandler) GetCurrentWeather(c *gin.Context) {
	location := c.Query("location")
	if err := utils.ValidateLocation(location); err != nil {
		utils.SendError(c, err)
		return
	}

	units := c.DefaultQuery("units", "metric")
	if err := utils.ValidateUnits(units); err != nil {
		utils.SendError(c, err)
		return
	}

	weatherData, err := h.weatherService.GetCurrentWeather(location, units)
	if err != nil {
		utils.SendError(c, utils.HandleWeatherAPIError(err))
		return
	}

	utils.SendSuccess(c, weatherData)
}

// GetWeatherForecast handles GET /weather/forecast requests
func (h *WeatherHandler) GetWeatherForecast(c *gin.Context) {
	location := c.Query("location")
	if err := utils.ValidateLocation(location); err != nil {
		utils.SendError(c, err)
		return
	}

	units := c.DefaultQuery("units", "metric")
	if err := utils.ValidateUnits(units); err != nil {
		utils.SendError(c, err)
		return
	}

	daysStr := c.DefaultQuery("days", "5")
	days, err := strconv.Atoi(daysStr)
	if err != nil {
		utils.SendError(c, utils.NewAPIError(http.StatusBadRequest, "Invalid days parameter", "Must be a valid number"))
		return
	}

	if err := utils.ValidateDays(days); err != nil {
		utils.SendError(c, err)
		return
	}

	weatherData, err := h.weatherService.GetWeatherForecast(location, units, days)
	if err != nil {
		utils.SendError(c, utils.HandleWeatherAPIError(err))
		return
	}

	utils.SendSuccess(c, weatherData)
}

// PostCurrentWeather handles POST /weather/current requests with JSON body
func (h *WeatherHandler) PostCurrentWeather(c *gin.Context) {
	var req models.WeatherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, utils.NewAPIError(http.StatusBadRequest, "Invalid request body", err.Error()))
		return
	}

	if err := utils.ValidateLocation(req.Location); err != nil {
		utils.SendError(c, err)
		return
	}

	units := req.Units
	if units == "" {
		units = "metric"
	}

	if err := utils.ValidateUnits(units); err != nil {
		utils.SendError(c, err)
		return
	}

	weatherData, err := h.weatherService.GetCurrentWeather(req.Location, units)
	if err != nil {
		utils.SendError(c, utils.HandleWeatherAPIError(err))
		return
	}

	utils.SendSuccess(c, weatherData)
}

// PostWeatherForecast handles POST /weather/forecast requests with JSON body
func (h *WeatherHandler) PostWeatherForecast(c *gin.Context) {
	var req models.WeatherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, utils.NewAPIError(http.StatusBadRequest, "Invalid request body", err.Error()))
		return
	}

	if err := utils.ValidateLocation(req.Location); err != nil {
		utils.SendError(c, err)
		return
	}

	units := req.Units
	if units == "" {
		units = "metric"
	}

	if err := utils.ValidateUnits(units); err != nil {
		utils.SendError(c, err)
		return
	}

	days := req.Days
	if days == 0 {
		days = 5
	}

	if err := utils.ValidateDays(days); err != nil {
		utils.SendError(c, err)
		return
	}

	weatherData, err := h.weatherService.GetWeatherForecast(req.Location, units, days)
	if err != nil {
		utils.SendError(c, utils.HandleWeatherAPIError(err))
		return
	}

	utils.SendSuccess(c, weatherData)
}

// HealthCheck handles GET /health requests
func (h *WeatherHandler) HealthCheck(c *gin.Context) {
	response := gin.H{
		"status":    "healthy",
		"service":   "weathering-with-go",
		"timestamp": gin.H{"unix": gin.H{}},
	}
	c.JSON(http.StatusOK, response)
}