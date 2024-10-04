package config

import (
	"fmt"
	"os"
)

// LoadConfig загружает конфигурацию из указанного файла
func LoadConfig(fileName string) map[string]string {
	fmt.Printf("Loading configuration from %s...\n", fileName)

	// В реальном проекте это будет чтение и парсинг файла конфигурации
	// Для простоты в данном примере используется заглушка
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("Configuration file not found!")
		return nil
	}

	// Пример конфигурации
	config := map[string]string{
		"device_ip": "192.168.1.1",
		"username":  "admin",
		"password":  "password",
	}

	return config
}
