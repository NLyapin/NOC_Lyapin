package main

import (
	"fmt"

	"github.com/NLyapin/NOC_Lyapin/internal/devices"
	"github.com/NLyapin/NOC_Lyapin/pkg/config"
)

func main() {
	fmt.Println("Network Configuration System is starting...")

	// Пример использования модуля конфигураций
	configData := config.LoadConfig("example-config.yaml")
	fmt.Println("Loaded configuration:", configData)

	// Пример взаимодействия с сетевым устройством
	device := devices.NewDevice("192.168.1.1", "admin", "password")
	status := device.Connect()
	if status {
		fmt.Println("Device connected successfully!")
	} else {
		fmt.Println("Failed to connect to the device.")
	}
}
