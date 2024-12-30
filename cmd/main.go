package main

import (
	"log"

	"config-manager/internal/client" // Импорт пакета client
	"config-manager/internal/gnmi"   // Импорт пакета gnmi
)

func main() {
	server := "192.0.2.1:57400"
	certFile := "certs/client.crt"
	keyFile := "certs/client.key"
	caFile := "certs/ca.crt"

	// Создание gNMI клиента
	conn, err := client.NewGNMIClient(server, certFile, keyFile, caFile) // Используем client.NewGNMIClient
	if err != nil {
		log.Fatalf("Ошибка при создании клиента: %v", err)
	}
	defer conn.Close()

	gnmiClient := gnmi.NewGNMIClient(conn) // Используем gnmi.NewGNMIClient для получения клиента

	// Чтение конфигурации интерфейса
	interfaceName := "eth0"
	config, err := gnmi.GetInterfaceConfig(gnmiClient, interfaceName)
	if err != nil {
		log.Fatalf("Ошибка чтения конфигурации интерфейса: %v", err)
	}

	log.Printf("Текущая конфигурация интерфейса %s: %v", interfaceName, config)

	// Обновление описания интерфейса
	newDescription := "Updated interface description"
	if err := gnmi.SetInterfaceDescription(gnmiClient, interfaceName, newDescription); err != nil {
		log.Fatalf("Ошибка обновления описания интерфейса: %v", err)
	}

	log.Println("Описание интерфейса успешно обновлено.")
}