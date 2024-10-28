package main

import (
	"fmt"
	"log"

	"github.com/NLyapin/NOC_Lyapin/internal/devices"
	"github.com/NLyapin/NOC_Lyapin/pkg/config"

	"golang.org/x/crypto/ssh"
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

func QoS() {
	// Параметры подключения SSH
	sshHost := "10.25.16.254:62214"
	sshUser := "user"
	sshPassword := "user6501"

	// Настройки SSH клиента
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Подключение к роутеру
	client, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		log.Fatalf("Ошибка подключения: %s", err)
	}
	defer client.Close()

	// Создание SSH сессии
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Ошибка создания сессии: %s", err)
	}
	defer session.Close()

	// Команды для настройки QoS
	command := `
    configure terminal
    qos profile name HighPriority
    qos priority high
    qos profile name LowPriority
    qos priority low
    interface ethernet 0/1
    qos apply profile HighPriority
    commit
    show qos configuration
    `

	// Выполнение команды и получение вывода
	output, err := session.CombinedOutput(command)
	if err != nil {
		log.Fatalf("Ошибка выполнения команды: %s", err)
	}

	// Вывод результата
	fmt.Println("Результат настройки QoS:")
	fmt.Println(string(output))
}
