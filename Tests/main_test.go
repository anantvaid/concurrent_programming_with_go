package main

import "testing"

func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	expectedArea := 50.0
	if area := rect.Area(); area != expectedArea {
		t.Errorf("Expected area %f, but got %f", expectedArea, area)
	}
}

func TestRectangleAreaTable(t *testing.T) {
	tests := []struct {
		name                        string
		width, height, expectedArea float64
	}{
		{"Test1", 10, 5, 50},
		{"Test2", 3, 4, 12},
		{"Test3", 7.5, 2, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rect := Rectangle{Width: tt.width, Height: tt.height}
			if area := rect.Area(); area != tt.expectedArea {
				t.Errorf("Expected area %f, but got %f", tt.expectedArea, area)
			}
		})
	}
}
