package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gosnmp/gosnmp"
	"golang.org/x/crypto/ssh"
)

// Параметры подключения
const (
	sshHost       = "10.25.16.254:62214"
	sshUser       = "user"
	sshPassword   = "user6501"
	snmpCommunity = "private"
	snmpTarget    = "192.168.1.1" // IP роутера
)

// QoS настройка в SNMP OID
var (
	qosOID = ".1.3.6.1.2.1.2.2.1.8" // Пример OID для QoS (замените на корректный для вашего устройства)
)

// Устанавливаем SSH-подключение и туннель к SNMP-серверу
func main() {
	fmt.Println("Старт программы")

	// Подключение по SSH
	fmt.Println("Устанавливаем SSH-подключение к промежуточному серверу...")
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		log.Fatalf("Ошибка подключения по SSH: %s", err)
	}
	defer client.Close()
	fmt.Println("SSH-подключение установлено")

	// Установка SNMP-клиента
	fmt.Println("Подключаемся к SNMP-серверу на устройстве...")
	snmp := &gosnmp.GoSNMP{
		Target:    snmpTarget,
		Port:      161,
		Community: snmpCommunity,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		Retries:   1,
	}

	err = snmp.Connect()
	if err != nil {
		log.Fatalf("Ошибка подключения к SNMP-серверу: %s", err)
	}
	defer snmp.Conn.Close()
	fmt.Println("SNMP-подключение успешно установлено")

	// Выполнение команды изменения QoS
	fmt.Println("Изменяем параметры QoS через SNMP...")
	value := 1 // Пример значения QoS (значение замените на корректное для вашего устройства)
	oidValue := gosnmp.SnmpPDU{
		Name:  qosOID,
		Type:  gosnmp.Integer,
		Value: value,
	}

	_, err = snmp.Set([]gosnmp.SnmpPDU{oidValue})
	if err != nil {
		log.Fatalf("Ошибка изменения QoS: %s", err)
	}
	fmt.Println("Параметры QoS успешно изменены")

	// Проверка изменений
	fmt.Println("Проверяем изменение QoS...")
	result, err := snmp.Get([]string{qosOID})
	if err != nil {
		log.Fatalf("Ошибка получения QoS: %s", err)
	}

	for _, variable := range result.Variables {
		fmt.Printf("Изменение QoS выполнено. Текущее значение: %d\n", variable.Value)
	}
	fmt.Println("Программа завершена успешно")
}
