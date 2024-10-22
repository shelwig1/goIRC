package main

import (
	"fmt"
	"net"
)

// Need to set a constraint on how bug the username can be
const username string = "shelwig"
const server_address = ":8080"

func main() {
	connect_to_server()
}

func connect_to_server() {
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

}

// I guess the point of this is to get the p2p shit working, so let's make it work here

// Request connection to the other guy, server holds the message in escrow
