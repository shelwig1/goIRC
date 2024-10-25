package main

import (
	"fmt"
	"net"
	"os"
)

type User struct {
	Username string
	Address  net.Addr
}

const (
	serverAddress = "localhost:8080"
)

var active_users []User

func main() {
	startServer()
}

// startServer starts a TCP server that listens for incoming connections.
func startServer() {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error listening on", serverAddress)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is listening on", serverAddress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

// handleConnection handles incoming connections.
func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)

	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			fmt.Println("Connection closed by user:", conn.RemoteAddr())
			return
		}

		packet := deserializePacket(buf[:n])
		routePacket(packet, conn)

	}
}

func createUser(packet Packet) User {
	user := User{Username: packet.Username, Address: packet.Address}
	return user
}

func initialConnection(packet Packet) {
	new_user := createUser(packet)

	active_users = append(active_users, new_user)

}
