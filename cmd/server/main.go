package main

import (
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:7777")

	if err != nil {
		log.Fatalln("could not create listener")
		return
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("could not handle connection")
		}

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	conn.Read(buf)
	log.Println(string(buf))

}


