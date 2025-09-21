package services

import (
	"testing"
	"time"

	"weathering-with-go/models"
)

func TestConvertCurrentWeatherResponse(t *testing.T) {
	svc := NewWeatherService("dummy")
	owm := models.OpenWeatherMapResponse{
		Coord:   models.Coordinates{Lat: 1.23, Lon: 4.56},
		Weather: []models.Weather{{Main: "Clear", Description: "clear sky", Icon: "01d"}},
		Main:    models.Main{Temp: 10.5, FeelsLike: 9.0, Pressure: 1012, Humidity: 80},
		Wind:    models.Wind{Speed: 3.4, Deg: 180},
		Clouds:  models.Clouds{All: 0},
		Dt:      time.Now().Unix(),
		Sys:     models.Sys{Country: "GB"},
		Name:    "Testville",
	}

	data := svc.convertCurrentWeatherResponse(owm)
	if data.Location.Name != "Testville" {
		t.Fatalf("expected location name Testville got %s", data.Location.Name)
	}
	if data.Current.Condition != "Clear" {
		t.Fatalf("expected condition Clear got %s", data.Current.Condition)
	}
}

func TestCalculateDailyForecastEmpty(t *testing.T) {
	svc := NewWeatherService("dummy")
	f := svc.calculateDailyForecast("2025-01-02", nil)
	if f.Date.IsZero() {
		t.Fatalf("expected non-zero date")
	}
}
