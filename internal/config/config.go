package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/openconfig/gnmi/client"
	"github.com/openconfig/gnmi/proto/gnmi"
)

func AddInterface(name, ifaceType string, enabled bool) error {
	log.Printf("⚙️ Добавление интерфейса %s...", name)

	update := &gnmi.SetRequest{
		Prefix: &gnmi.Path{Target: "router1"},
		Update: []*gnmi.Update{
			{
				Path: &gnmi.Path{Elem: []*gnmi.PathElem{{Name: "interfaces"}, {Name: "interface", Key: map[string]string{"name": name}}}},
				Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{
					StringVal: fmt.Sprintf(`{ "config": { "name": "%s", "type": "%s", "enabled": %t } }`, name, ifaceType, enabled),
				}},
			},
		},
	}

	return sendConfig(update)
}

func DeleteRoute(prefix string) error {
	log.Printf("❌ Удаление маршрута %s...", prefix)

	delete := &gnmi.SetRequest{
		Prefix: &gnmi.Path{Target: "router1"},
		Delete: []*gnmi.Path{
			{Elem: []*gnmi.PathElem{
				{Name: "network-instances"},
				{Name: "network-instance", Key: map[string]string{"name": "default"}},
				{Name: "table-connections"},
				{Name: "static"},
				{Name: "routes"},
				{Name: "route", Key: map[string]string{"prefix": prefix}},
			}},
		},
	}

	return sendConfig(delete)
}

func sendConfig(req *gnmi.SetRequest) error {
	addr := "localhost:50051"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	target := client.Destination{
		Addr:     addr,
		Target:   "router1",
		Encoding: gnmi.Encoding_JSON,
	}

	cli, err := client.New(ctx, client.Type("gnmi"), target)
	if err != nil {
		return fmt.Errorf("Ошибка подключения: %v", err)
	}
	defer cli.Close()

	resp, err := cli.(*client.Impl).Set(ctx, req)
	if err != nil {
		return fmt.Errorf("Ошибка отправки: %v", err)
	}

	log.Printf("✅ Конфигурация обновлена: %v", resp)
	return nil
}

func AddRoute(prefix, nexthop string) error {
	log.Printf("🛣 Добавление маршрута %s через %s...", prefix, nexthop)

	update := &gnmi.SetRequest{
		Prefix: &gnmi.Path{Target: "router1"},
		Update: []*gnmi.Update{
			{
				Path: &gnmi.Path{Elem: []*gnmi.PathElem{
					{Name: "network-instances"},
					{Name: "network-instance", Key: map[string]string{"name": "default"}},
					{Name: "table-connections"},
					{Name: "static"},
					{Name: "routes"},
					{Name: "route", Key: map[string]string{"prefix": prefix}},
				}},
				Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{
					StringVal: fmt.Sprintf(`{ "config": { "prefix": "%s", "nexthop": "%s" } }`, prefix, nexthop),
				}},
			},
		},
	}

	return sendConfig(update)
}

func DeleteInterface(name string) error {
	log.Printf("❌ Удаление интерфейса %s...", name)

	delete := &gnmi.SetRequest{
		Prefix: &gnmi.Path{Target: "router1"},
		Delete: []*gnmi.Path{
			{Elem: []*gnmi.PathElem{{Name: "interfaces"}, {Name: "interface", Key: map[string]string{"name": name}}}},
		},
	}

	return sendConfig(delete)
}

func SetIPAddress(interfaceName, ipAddress string) error {
	log.Printf("🌍 Установка IP %s на интерфейсе %s...", ipAddress, interfaceName)

	update := &gnmi.SetRequest{
		Prefix: &gnmi.Path{Target: "router1"},
		Update: []*gnmi.Update{
			{
				Path: &gnmi.Path{Elem: []*gnmi.PathElem{
					{Name: "interfaces"},
					{Name: "interface", Key: map[string]string{"name": interfaceName}},
					{Name: "subinterfaces"},
					{Name: "subinterface", Key: map[string]string{"index": "0"}},
					{Name: "ipv4"},
					{Name: "addresses"},
					{Name: "address", Key: map[string]string{"ip": ipAddress}},
				}},
				Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{
					StringVal: fmt.Sprintf(`{ "config": { "ip": "%s", "prefix-length": 24 } }`, ipAddress),
				}},
			},
		},
	}

	return sendConfig(update)
}

// Функция для обновления конфигурации маршрутизатора
func UpdateConfig() error {
	log.Println("⚙️ Изменение конфигурации маршрутизатора...")

	addr := "localhost:50051"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	target := client.Destination{
		Addr:     addr,
		Target:   "router1",
		Encoding: gnmi.Encoding_JSON,
	}

	cli, err := client.New(ctx, client.Type("gnmi"), target)
	if err != nil {
		return fmt.Errorf("❌ Ошибка подключения: %v", err)
	}
	defer cli.Close()

	// Пример: добавляем новый интерфейс Loopback1
	update := &gnmi.SetRequest{
		Prefix: &gnmi.Path{Target: target.Target},
		Update: []*gnmi.Update{
			{
				Path: &gnmi.Path{Elem: []*gnmi.PathElem{{Name: "interfaces"}, {Name: "interface", Key: map[string]string{"name": "Loopback1"}}}},
				Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{StringVal: `{ "config": { "name": "Loopback1", "type": "iana-if-type:softwareLoopback", "enabled": true } }`}},
			},
		},
	}

	resp, err := cli.(*client.Impl).Set(ctx, update)
	if err != nil {
		return fmt.Errorf("❌ Ошибка обновления: %v", err)
	}

	log.Printf("✅ Конфигурация изменена: %v", resp)
	return nil
}
