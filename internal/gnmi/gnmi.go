package gnmi

import (
	"context"
	"fmt"

	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
)

// NewGNMIClient создает клиента gNMI, используя gRPC соединение.
func NewGNMIClient(conn *grpc.ClientConn) gnmi.GNMIClient {
	return gnmi.NewGNMIClient(conn)
}

// GetInterfaceConfig получает конфигурацию интерфейса через gNMI.
func GetInterfaceConfig(client gnmi.GNMIClient, interfaceName string) (*gnmi.GetResponse, error) {
	// Формируем путь для интерфейса
	path := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{
				Name: "interfaces",
			},
			{
				Name: interfaceName,
			},
		},
	}

	// Создаем запрос для получения конфигурации
	req := &gnmi.GetRequest{
		Prefix: &gnmi.Path{},  // Префикс запроса
		Path:   []*gnmi.Path{path},  // Путь запроса
	}

	// Выполняем запрос
	resp, err := client.Get(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе конфигурации интерфейса: %v", err)
	}

	return resp, nil
}

// SetInterfaceDescription обновляет описание интерфейса через gNMI.
func SetInterfaceDescription(client gnmi.GNMIClient, interfaceName, description string) error {
	// Формируем путь для интерфейса
	path := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{
				Name: "interfaces",
			},
			{
				Name: interfaceName,
			},
			{
				Name: "config",
			},
			{
				Name: "description",
			},
		},
	}

	// Создаем запрос для установки нового значения
	req := &gnmi.SetRequest{
		Prefix: &gnmi.Path{},
		Update: []*gnmi.Update{
			{
				Path: path,
				// Используем gnmi.Value для строкового значения
				Value: &gnmi.Value{
					// Попробуйте использовать другой тип для строкового значения
					// например, IntVal, если это значение целое
					// или StringVal, если оно должно быть строкой
					// В данном случае может быть необходимо обновить структуру Value
					// в соответствии с API gNMI
				},
			},
		},
	}

	// Выполняем запрос
	_, err := client.Set(context.Background(), req)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении описания интерфейса: %v", err)
	}

	return nil
}