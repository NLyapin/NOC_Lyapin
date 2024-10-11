package config

import (
	"fmt"
	"sync"
)

// ConfigStore хранит конфигурации в памяти
type ConfigStore struct {
	configs map[string]map[string]string
	mu      sync.RWMutex
}

// NewConfigStore создает новый экземпляр ConfigStore
func NewConfigStore() *ConfigStore {
	return &ConfigStore{
		configs: make(map[string]map[string]string),
	}
}

// AddConfig добавляет новую конфигурацию
func (cs *ConfigStore) AddConfig(name string, config map[string]string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.configs[name] = config
	fmt.Printf("Added configuration: %s\n", name)
}

// UpdateConfig обновляет существующую конфигурацию
func (cs *ConfigStore) UpdateConfig(name string, config map[string]string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	if _, exists := cs.configs[name]; exists {
		cs.configs[name] = config
		fmt.Printf("Updated configuration: %s\n", name)
	} else {
		fmt.Printf("Configuration %s does not exist.\n", name)
	}
}

// DeleteConfig удаляет конфигурацию
func (cs *ConfigStore) DeleteConfig(name string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	delete(cs.configs, name)
	fmt.Printf("Deleted configuration: %s\n", name)
}

// GetConfig возвращает конфигурацию по имени
func (cs *ConfigStore) GetConfig(name string) (map[string]string, bool) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	config, exists := cs.configs[name]
	return config, exists
}
