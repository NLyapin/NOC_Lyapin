package main

import (
	"log"

	"config-manager/internal/client" // Используем client для создания gNMI клиента
	"config-manager/internal/gnmi"   // Пакет gnmi для работы с gNMI
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

	// Создаем gNMI клиента с помощью gRPC соединения
	gnmiClient := gnmi.NewGNMIClient(conn) // Это вызов функции из пакета gnmi, которая использует grpc соединение

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