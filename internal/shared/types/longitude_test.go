package types

import (
	"testing"
)

func TestLongitude_IsValid_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"0", true},
		{"45.123", true},
		{"-180", true},
		{"180", true},
		{"-73.935242", true},
	}

	for _, test := range tests {
		l := Longitude{Value: test.input}
		valid, err := l.IsValid()
		if valid != test.expected || err != nil {
			t.Errorf("Expected %v, got %v with error %v for input %s", test.expected, valid, err, test.input)
		}
	}
}

func TestLongitude_IsValid_InvalidRange(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"-190.5"},
		{"200"},
		{"181"},
		{"-181"},
	}

	for _, test := range tests {
		l := Longitude{Value: test.input}
		valid, err := l.IsValid()
		if valid || err == nil {
			t.Errorf("Expected error for input %s, but got valid=%v, err=%v", test.input, valid, err)
		}
	}
}

func TestLongitude_IsValid_NotANumber(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"abc"},
		{"123abc"},
		{"--45"},
		{"1.2.3"},
		{" "},
	}

	for _, test := range tests {
		l := Longitude{Value: test.input}
		valid, err := l.IsValid()
		if valid || err == nil {
			t.Errorf("Expected error for input %s, but got valid=%v, err=%v", test.input, valid, err)
		}
	}
}

func TestLongitude_IsEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{" ", false},
		{"0", false},
		{"45.123", false},
	}

	for _, test := range tests {
		l := Longitude{Value: test.input}
		isEmpty := l.IsEmpty()
		if isEmpty != test.expected {
			t.Errorf("Expected IsEmpty() to return %v for input %s, got %v", test.expected, test.input, isEmpty)
		}
	}
}

