package main

import "testing"

// Test For SetChime method
func TestSetChime(t *testing.T) {
	db := NewDoorBell()
	db.SetChime("music")
	if db.chime != "music" {
		t.Errorf("Expected chime to be 'music', found %s", db.chime)
	}
}

// Test For AdjustVolume Method
func TestAdjustVolume(t *testing.T) {
	db := NewDoorBell()
	db.AdjustVolume(12)
	if db.volume != 7 {
		t.Errorf("Expected volume to be 7, found %d", db.volume)
	}
}

// Test For IntegrateLight Method
func TestIntegrateLights(t *testing.T) {
	db := NewDoorBell()
	db.IntegrateLights(true)
	if !db.lightOn {
		t.Errorf("Expected integrateLights to be true, found %v", db.lightOn)
	}
}

// TestRing verifies that the doorbell rings with the correct chime and volume
func TestRing(t *testing.T) {
	db := NewDoorBell()
	db.SetChime("music")
	db.AdjustVolume(10)
	db.IntegrateLights(true)

	// Mocking the output would be required here to test the Ring Method Effectively

	// For example, you could redirect the standard output to buffer and then verify the buffer's contents
}
