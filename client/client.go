package main

import (
	"fmt"
	"net"
)

// Need to set a constraint on how bug the username can be
const username string = "shelwig"
const server_address = ":8080"

func main() {
	shutdown := make(chan struct{})
	connectToServer(shutdown)
}

/* func connectToServer() {
	conn, err := net.Dial("tcp", server_address)
	if err != nil {
		fmt.Println("Error dialing server")
	}

	defer conn.Close()

	conn.Write([]byte(username))

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			continue
		} else {
			fmt.Println(string(buf[:n]))

		}

	}
} */

func connectToServer(shutdown chan struct{}) {
	conn, err := net.Dial("tcp", server_address)
	if err != nil {
		fmt.Println("Error dialing server")
		return
	}
	defer conn.Close()

	conn.Write([]byte(username))

	buf := make([]byte, 1024)

	for {
		select {
		case <-shutdown:
			fmt.Println("Shutdown signal received. Exiting...")
			return // Exit on shutdown signal
		default:
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Error reading from server:", err)
				return // Exit on read error
			} else {
				fmt.Println(string(buf[:n]))
			}
		}
	}
}

/* func newConnect() {
	conn, err := net.Dial("tcp", server_address)
	if err != nil {
		fmt.Println("Error dialing server")
	}

	defer conn.Close()

	conn.Write(initialConnection())

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			continue
		} else {
			fmt.Println(string(buf[:n]))
			// Route the packet appropriately
			packet := readRawPacket(buf[:n])
			routePacket(packet)
		}

	}
}

func routePacket(packet Packet) {

}

func initialConnection() []byte {
	packet := Packet{User: username, Goal: "initial_connection"}
	data, err := json.Marshal(packet)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func readRawPacket(data []byte) Packet {
	var packet Packet
	err := json.Unmarshal(data, &packet)
	if err != nil {
		log.Fatal(err)
	}

	return packet
}
*/
// I guess the point of this is to get the p2p shit working, so let's make it work here

// Request connection to the other guy, server holds the message in escrow
