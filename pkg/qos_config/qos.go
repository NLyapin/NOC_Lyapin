package main

import (
	"log"
	"os/exec"

	"golang.org/x/crypto/ssh"
)

func main() {
	QoS() // Вызов функции QoS
}

func QoS() {
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

	log.Println("Подключение к устройству...")
	client, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		log.Fatalf("Ошибка подключения: %s", err)
	}
	defer client.Close()
	log.Println("Подключение успешно!")

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Ошибка создания сессии: %s", err)
	}
	defer session.Close()

	// Запуск minicom
	minicomCommand := exec.Command("minicom", "-wD", "/dev/usbports/st10_esr")
	if err := minicomCommand.Start(); err != nil {
		log.Fatalf("Ошибка запуска minicom: %s", err)
	}

	// Здесь предполагается, что вы входите в minicom, используя учетные данные admin/password
	log.Println("Запуск minicom... Войдите как 'admin' с паролем 'password'.")

	// Подождите, пока minicom завершится
	if err := minicomCommand.Wait(); err != nil {
		log.Fatalf("Ошибка minicom: %s", err)
	}

	// Если требуется, выполните дополнительные команды после выхода из minicom
	log.Println("Выход из minicom. Продолжение выполнения программы.")
}
