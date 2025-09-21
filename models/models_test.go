package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAPIResponseMarshal(t *testing.T) {
	resp := APIResponse{
		Success: true,
		Data:    map[string]interface{}{"foo": "bar"},
	}

	b, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var out APIResponse
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if !out.Success {
		t.Fatalf("expected success true")
	}
}

func TestWeatherRequestTags(t *testing.T) {
	rt := reflect.TypeOf(WeatherRequest{})
	f, ok := rt.FieldByName("Location")
	if !ok {
		t.Fatalf("Location field missing")
	}
	jsonTag := f.Tag.Get("json")
	formTag := f.Tag.Get("form")
	if jsonTag != "location" || formTag != "location" {
		t.Fatalf("unexpected tags: json=%s form=%s", jsonTag, formTag)
	}
}
