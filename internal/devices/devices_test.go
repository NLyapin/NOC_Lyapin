package devices

import (
	"testing"
)

func TestDevice(t *testing.T) {
	device := NewDevice("192.168.1.1", "admin", "password")

	if !device.Connect() {
		t.Fatal("Failed to connect to the device")
	}

	info := device.GetDeviceInfo()
	if info != "Device IP: 192.168.1.1, Username: admin" {
		t.Fatalf("Device info mismatch: got %s", info)
	}
}
