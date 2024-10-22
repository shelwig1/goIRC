package main

import (
	"fmt"
	"net"
)

// Need to set a constraint on how bug the username can be
const username string = "shelwig"
const server_address = ":8080"

/*
	 func main() {
		// dial to the main server

		// receive the list of connections

}
*/
func main() {
	connect_to_server()
}

func connect_to_server() {
	conn, err := net.Dial("tcp", server_address)
	if err != nil {
		fmt.Println("Error dialing server")
	}

	// When we establish the connection, send our goodies

	// username
	conn.Write([]byte(username))

}
