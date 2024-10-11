package config

import (
	"testing"
)

func TestConfigStore(t *testing.T) {
	cs := NewConfigStore()

	// Тестирование добавления конфигурации
	cs.AddConfig("TestRouter", map[string]string{"key": "value"})
	config, exists := cs.GetConfig("TestRouter")
	if !exists || config["key"] != "value" {
		t.Fatalf("Failed to add or retrieve configuration")
	}

	// Тестирование обновления конфигурации
	cs.UpdateConfig("TestRouter", map[string]string{"key": "newvalue"})
	config, _ = cs.GetConfig("TestRouter")
	if config["key"] != "newvalue" {
		t.Fatalf("Failed to update configuration")
	}

	// Тестирование удаления конфигурации
	cs.DeleteConfig("TestRouter")
	_, exists = cs.GetConfig("TestRouter")
	if exists {
		t.Fatalf("Failed to delete configuration")
	}
}
