package devices

import (
	"fmt"
)

// Device представляет сетевое устройство
type Device struct {
	IP       string
	Username string
	Password string
}

// NewDevice создает новый экземпляр сетевого устройства
func NewDevice(ip, username, password string) *Device {
	return &Device{
		IP:       ip,
		Username: username,
		Password: password,
	}
}

// Connect выполняет подключение к устройству
func (d *Device) Connect() bool {
	fmt.Printf("Connecting to device at %s with user %s...\n", d.IP, d.Username)
	// Заглушка для успешного подключения
	return true
}

// GetDeviceInfo получает информацию о устройстве (заглушка)
func (d *Device) GetDeviceInfo() string {
	return fmt.Sprintf("Device IP: %s, Username: %s", d.IP, d.Username)
}
