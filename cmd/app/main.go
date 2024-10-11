package main

import (
	"fmt"

	"github.com/NLyapin/NOC_Lyapin/internal/devices"
	"github.com/NLyapin/NOC_Lyapin/pkg/config"
)

func main() {
	fmt.Println("Network Configuration System is starting...")

	// Создаем хранилище конфигураций
	configStore := config.NewConfigStore()

	// Пример добавления конфигурации
	configStore.AddConfig("Router1", map[string]string{
		"device_ip": "192.168.1.1",
		"username":  "admin",
		"password":  "password",
	})

	// Получение конфигурации
	if config, exists := configStore.GetConfig("Router1"); exists {
		fmt.Println("Loaded configuration:", config)
	}

	// Пример взаимодействия с сетевым устройством
	device := devices.NewDevice("192.168.1.1", "admin", "password")
	if device.Connect() {
		fmt.Println("Device connected successfully!")
		fmt.Println(device.GetDeviceInfo())
	} else {
		fmt.Println("Failed to connect to the device.")
	}

	// Обновление конфигурации
	configStore.UpdateConfig("Router1", map[string]string{
		"device_ip": "192.168.1.2",
		"username":  "admin",
		"password":  "newpassword",
	})

	// Удаление конфигурации
	configStore.DeleteConfig("Router1")
}
