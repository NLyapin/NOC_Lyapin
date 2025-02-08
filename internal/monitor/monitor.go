package monitor

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/openconfig/gnmi/client"
	"github.com/openconfig/gnmi/proto/gnmi"
)

// Получение конфигурации маршрутизатора через gNMI
func GetConfig() error {
	log.Println("📡 Получение конфигурации маршрутизатора...")

	addr := "localhost:50051"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	target := client.Destination{
		Addr:     addr,
		Target:   "router1",
		Encoding: gnmi.Encoding_JSON,
	}

	// Создаём gNMI клиент
	cli, err := client.New(ctx, client.Type("gnmi"), target)
	if err != nil {
		return fmt.Errorf("❌ Ошибка подключения: %v", err)
	}
	defer cli.Close()

	// Запрос текущей конфигурации интерфейсов
	req := &gnmi.GetRequest{
		Prefix: &gnmi.Path{Target: target.Target},
		Path: []*gnmi.Path{
			{Elem: []*gnmi.PathElem{{Name: "interfaces"}}},
		},
		Encoding: gnmi.Encoding_JSON,
	}

	resp, err := cli.(*client.Impl).Get(ctx, req)
	if err != nil {
		return fmt.Errorf("❌ Ошибка получения данных: %v", err)
	}

	log.Printf("✅ Текущая конфигурация: %v", resp)
	return nil
}
