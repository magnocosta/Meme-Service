package types

import (
	"testing"
)

func TestCoordinate_IsValid_ValidValues(t *testing.T) {
	tests := []struct {
		latInput string
		lonInput string
		expected bool
	}{
		{"40.7128", "-74.0060", true}, // Example: New York
		{"-33.8651", "151.2099", true}, // Example: Sydney
		{"0", "0", true},               // Edge case: Equator & Prime Meridian
		{"90", "180", true},            // Edge case: Max valid values
		{"-90", "-180", true},          // Edge case: Min valid values
	}

	for _, test := range tests {
		c := Coordinate{
			Lat: Latitude{Value: test.latInput},
			Lon: Longitude{Value: test.lonInput},
		}
		valid, err := c.IsValid()
		if valid != test.expected || err != nil {
			t.Errorf("Expected %v, got %v with error %v for input (Lat: %s, Lon: %s)", test.expected, valid, err, test.latInput, test.lonInput)
		}
	}
}

func TestCoordinate_IsValid_InvalidLongitude(t *testing.T) {
	tests := []struct {
		latInput string
		lonInput string
	}{
		{"40.7128", "-190"},
		{"-33.8651", "200"},
		{"0", "181"},
	}

	for _, test := range tests {
		c := Coordinate{
			Lat: Latitude{Value: test.latInput},
			Lon: Longitude{Value: test.lonInput},
		}
		valid, err := c.IsValid()
		if valid || err == nil {
			t.Errorf("Expected error for input (Lat: %s, Lon: %s), but got valid=%v, err=%v", test.latInput, test.lonInput, valid, err)
		}
	}
}

func TestCoordinate_IsValid_InvalidLatitude(t *testing.T) {
	tests := []struct {
		latInput string
		lonInput string
	}{
		{"100", "-74.0060"},
		{"-95", "151.2099"},
		{"91", "0"},
	}

	for _, test := range tests {
		c := Coordinate{
			Lat: Latitude{Value: test.latInput},
			Lon: Longitude{Value: test.lonInput},
		}
		valid, err := c.IsValid()
		if valid || err == nil {
			t.Errorf("Expected error for input (Lat: %s, Lon: %s), but got valid=%v, err=%v", test.latInput, test.lonInput, valid, err)
		}
	}
}

func TestCoordinate_IsValid_InvalidBoth(t *testing.T) {
	tests := []struct {
		latInput string
		lonInput string
	}{
		{"200", "200"},
		{"-200", "-190"},
	}

	for _, test := range tests {
		c := Coordinate{
			Lat: Latitude{Value: test.latInput},
			Lon: Longitude{Value: test.lonInput},
		}
		valid, err := c.IsValid()
		if valid || err == nil {
			t.Errorf("Expected error for input (Lat: %s, Lon: %s), but got valid=%v, err=%v", test.latInput, test.lonInput, valid, err)
		}
	}
}

func TestCoordinate_IsEmpty(t *testing.T) {
	tests := []struct {
		latInput  string
		lonInput  string
		expected  bool
	}{
		{"", "", true},
		{" ", " ", false},
		{"40.7128", "", false},
		{"", "-74.0060", false},
		{"0", "0", false},
	}

	for _, test := range tests {
		c := Coordinate{
			Lat: Latitude{Value: test.latInput},
			Lon: Longitude{Value: test.lonInput},
		}
		isEmpty := c.IsEmpty()
		if isEmpty != test.expected {
			t.Errorf("Expected IsEmpty() to return %v for input (Lat: %s, Lon: %s), got %v", test.expected, test.latInput, test.lonInput, isEmpty)
		}
	}
}

