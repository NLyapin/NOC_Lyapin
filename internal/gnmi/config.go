package gnmi

import (
	"context"
	"fmt"
	"time"

	"github.com/openconfig/gnmi/proto/gnmi"
)

// GetConfig выполняет gNMI Get запрос.
func GetConfig(client gnmi.GNMIClient, path *gnmi.Path) (*gnmi.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &gnmi.GetRequest{
		Path: []*gnmi.Path{path},
		Type: gnmi.GetRequest_CONFIG,
	}

	resp, err := client.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения Get запроса: %v", err)
	}

	return resp, nil
}

// SetConfig выполняет gNMI Set запрос.
func SetConfig(client gnmi.GNMIClient, updates []*gnmi.Update) (*gnmi.SetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &gnmi.SetRequest{
		Update: updates,
	}

	resp, err := client.Set(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения Set запроса: %v", err)
	}

	return resp, nil
}