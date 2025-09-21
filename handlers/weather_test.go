package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"weathering-with-go/services"

	"github.com/gin-gonic/gin"
)

// transportRedirect rewrites requests to point to the test server URL
type transportRedirect struct {
	target string
}

func (t *transportRedirect) RoundTrip(req *http.Request) (*http.Response, error) {
	// clone request
	r2 := new(http.Request)
	*r2 = *req
	// rewrite to target
	if strings.HasPrefix(r2.URL.String(), services.OpenWeatherMapBaseURL) || r2.URL.Host != "" {
		// replace host+scheme with target
		target := strings.TrimPrefix(t.target, "http://")
		r2.URL.Scheme = "http"
		r2.URL.Host = target
	}
	return http.DefaultTransport.RoundTrip(r2)
}

func TestGetCurrentWeatherHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// start a test server to simulate OpenWeatherMap
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"coord":{"lon":4.56,"lat":1.23},"weather":[{"main":"Clear","description":"clear sky","icon":"01d"}],"main":{"temp":10.5,"feels_like":9,"pressure":1012,"humidity":80},"wind":{"speed":3.4,"deg":180},"clouds":{"all":0},"dt":1234567890,"sys":{"country":"GB"},"name":"Testville","cod":200}`)
	}))
	defer srv.Close()

	svc := services.NewWeatherService("dummy")
	svc.HTTPClient = &http.Client{Transport: &transportRedirect{target: srv.URL}}

	wh := NewWeatherHandler(svc)
	router.GET("/api/v1/weather/current", wh.GetCurrentWeather)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/weather/current?location=Testville&units=metric", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK got %d body=%s", w.Code, w.Body.String())
	}
}
