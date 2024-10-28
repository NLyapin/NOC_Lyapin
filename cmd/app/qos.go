package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func QoS_test() {
	// Параметры подключения SSH
	sshHost := "10.25.16.254:62214"
	sshUser := "user"
	sshPassword := "user6501"

	// Настройки SSH клиента
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Подключение к роутеру
	client, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		log.Fatalf("Ошибка подключения: %s", err)
	}
	defer client.Close()

	// Создание SSH сессии
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Ошибка создания сессии: %s", err)
	}
	defer session.Close()

	// Команды для настройки QoS
	command := `
    configure terminal
    qos profile name HighPriority
    qos priority high
    qos profile name LowPriority
    qos priority low
    interface ethernet 0/1
    qos apply profile HighPriority
    commit
    show qos configuration
    `

	// Выполнение команды и получение вывода
	output, err := session.CombinedOutput(command)
	if err != nil {
		log.Fatalf("Ошибка выполнения команды: %s", err)
	}

	// Вывод результата
	fmt.Println("Результат настройки QoS:")
	fmt.Println(string(output))
}
