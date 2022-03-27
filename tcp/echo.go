package tcp

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

type EchoHandler struct {
	activeConn sync.Map
	closing    Boolean
}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

type EchoClient struct {
	Conn    net.Conn
	Waiting Wait
}

func (c *EchoClient) Close() error {
	c.Waiting.WaitWithTimeout(10 * time.Second)
	c.Conn.Close()
	return nil
}

func (h *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if h.closing.Get() {
		_ = conn.Close()
	}

	client := &EchoClient{
		Conn: conn,
	}
	h.activeConn.Store(client, struct{}{})

	reader := bufio.NewReader(client.Conn)
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
		client.Conn.Write(b)
	}
}

func (h *EchoHandler) Close() error {
	log.Println("handler shutting down...")
	h.closing.Set(true)
	h.activeConn.Range(func(key, value interface{}) bool {
		client := key.(*EchoClient)
		_ = client.Close()
		return true
	})
	return nil
}
