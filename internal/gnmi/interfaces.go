package gnmi

import (
	"fmt"

	"github.com/openconfig/gnmi/proto/gnmi"
)

// GetInterfaceConfig получает конфигурацию интерфейса.
func GetInterfaceConfig(client gnmi.GNMIClient, name string) (*gnmi.GetResponse, error) {
	path := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{Name: "interfaces"},
			{Name: "interface", Key: map[string]string{"name": name}},
			{Name: "config"},
		},
	}
	return GetConfig(client, path)
}

// SetInterfaceDescription обновляет описание интерфейса.
func SetInterfaceDescription(client gnmi.GNMIClient, name, description string) error {
	path := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{Name: "interfaces"},
			{Name: "interface", Key: map[string]string{"name": name}},
			{Name: "config"},
			{Name: "description"},
		},
	}
	update := &gnmi.Update{
		Path: path,
		Val: &gnmi.TypedValue{
			Value: &gnmi.TypedValue_StringVal{StringVal: description},
		},
	}

	_, err := SetConfig(client, []*gnmi.Update{update})
	if err != nil {
		return fmt.Errorf("ошибка обновления описания интерфейса: %v", err)
	}
	return nil
}