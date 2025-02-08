package main

import (
	"log"
	"openconfig-monitor/internal/config"
	"openconfig-monitor/internal/monitor"
	"os"
	"time"
)

func main() {
	log.Println("🚀 Запуск OpenConfig Monitor...")

	// Если переданы аргументы, выполняем команду
	if len(os.Args) > 1 {
		command := os.Args[1]

		switch command {
		case "add-interface":
			err := config.AddInterface("Loopback2", "iana-if-type:softwareLoopback", true)
			if err != nil {
				log.Fatalf("Ошибка добавления интерфейса: %v", err)
			}
		case "set-ip":
			if len(os.Args) < 4 {
				log.Fatal("Использование: go run cmd/main.go set-ip <interface> <ip>")
			}
			err := config.SetIPAddress(os.Args[2], os.Args[3])
			if err != nil {
				log.Fatalf("Ошибка установки IP: %v", err)
			}
		case "delete-interface":
			if len(os.Args) < 3 {
				log.Fatal("Использование: go run cmd/main.go delete-interface <interface>")
			}
			err := config.DeleteInterface(os.Args[2])
			if err != nil {
				log.Fatalf("Ошибка удаления интерфейса: %v", err)
			}
		case "add-route":
			if len(os.Args) < 4 {
				log.Fatal("Использование: go run cmd/main.go add-route <prefix> <nexthop>")
			}
			err := config.AddRoute(os.Args[2], os.Args[3])
			if err != nil {
				log.Fatalf("Ошибка добавления маршрута: %v", err)
			}
		case "delete-route":
			if len(os.Args) < 3 {
				log.Fatal("Использование: go run cmd/main.go delete-route <prefix>")
			}
			err := config.DeleteRoute(os.Args[2])
			if err != nil {
				log.Fatalf("Ошибка удаления маршрута: %v", err)
			}
		default:
			log.Fatal("Неизвестная команда")
		}

		return
	}

	// Если команд нет, просто мониторим состояние
	go func() {
		for {
			err := monitor.GetConfig()
			if err != nil {
				log.Printf("Ошибка мониторинга: %v", err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	select {} // Бесконечное ожидание (чтобы программа не завершалась)
}
