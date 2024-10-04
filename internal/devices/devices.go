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

	// В реальном проекте это будет подключение через SSH или другой протокол
	// Здесь представлена заглушка
	if d.IP == "" || d.Username == "" || d.Password == "" {
		fmt.Println("Invalid device credentials")
		return false
	}

	// Если всё в порядке, возвращаем true как успешное подключение
	return true
}
