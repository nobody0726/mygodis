package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func ListenAndServe(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("listen err %v", err))
	}
	defer listener.Close()
	log.Println(fmt.Sprintf("binding address %s, start listening...", address))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("accept error: %v", err))
			return
		}
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			return
		}
		b := []byte(msg)
		conn.Write(b)
	}
}

func main() {
	ListenAndServe(":8000")
}
