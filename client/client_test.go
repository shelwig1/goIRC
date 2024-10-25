package main

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func startMockServer(t *testing.T) {
	listener, err := net.Listen("tcp", ":8080")

	//listener, err := net.Listen("tcp", server_address)
	fmt.Println("Test: started mock server")
	if err != nil {
		t.Fatalf("Error starting mock server: %v", err)
	}

	defer listener.Close()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Logf("Error accepting connection: %v", err)
				continue
			}

			_, err = conn.Write([]byte("Hello from the mock server!"))
			if err != nil {
				t.Logf("Error writing to the connection: %v", err)
			}

			conn.Close()
		}
	}()
}

func TestConnectToServer(t *testing.T) {
	startMockServer(t)

	shutdown := make(chan struct{})

	time.Sleep(5 * time.Second)

	go connectToServer(shutdown)

	time.Sleep(5 * time.Second)
	shutdown <- struct{}{}

	time.Sleep(1 * time.Second)
}
