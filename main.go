package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	localAddr = "0.0.0.0:12345" // адрес запуска сервера
	proto     = "tcp4"          // протокол подключения
)

func main() {
	// Мапа поговорок с сайта
	proverbs := getProverbs()
	// Создаем объект listener сервера
	listener, err := net.Listen(proto, localAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		// Принимаем все входящие подключения
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Обрабатываем каждое новое соединение в отдельной горутине
		go handleConn(conn, proverbs)
	}
}

func handleConn(conn net.Conn, proverbs map[int]string) {
	defer conn.Close()

	for {
		rand.Seed(time.Now().UnixNano())
		//Случайный int для выборки из map поговорок
		randMapID := rand.Intn(len(proverbs))
		// Условие на случай если рандомизатор выдаст меньше 0
		if randMapID <= 0 {
			randMapID = 1
		}
		// Записываем случайный элемент из массива и переносим каретку на следующую строку
		conn.Write([]byte(proverbs[randMapID] + "\r\n"))
		// Ожидание 3 секунды до следующей итерации
		time.Sleep(3 * time.Second)
	}
}
