package utils

import "testing"

func TestValidateLocation(t *testing.T) {
	if err := ValidateLocation(""); err == nil {
		t.Fatalf("expected error for empty location")
	}
	if err := ValidateLocation("A"); err == nil {
		t.Fatalf("expected error for too short location")
	}
	if err := ValidateLocation("London,UK"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateUnits(t *testing.T) {
	if err := ValidateUnits(""); err != nil {
		t.Fatalf("empty units should be allowed")
	}
	if err := ValidateUnits("metric"); err != nil {
		t.Fatalf("metric should be valid")
	}
	if err := ValidateUnits("invalid"); err == nil {
		t.Fatalf("expected error for invalid units")
	}
}

func TestValidateDays(t *testing.T) {
	if err := ValidateDays(0); err == nil {
		t.Fatalf("expected error for days=0")
	}
	if err := ValidateDays(6); err == nil {
		t.Fatalf("expected error for days>5")
	}
	if err := ValidateDays(3); err != nil {
		t.Fatalf("unexpected error for days=3: %v", err)
	}
}
