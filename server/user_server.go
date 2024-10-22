package main

import (
	"fmt"
	"log"
	"net"
)

const server_address string = ":8080"

type User struct {
	Name    string
	Address net.Addr
}

var active_users []User

func main() {
	// listen on whatever dealio

	// whenever we get a client, spin it off and list off all the connections we need

	// What data gets sent with the initial connection?
	// Let's go with username required on initial connection
	listener, err := net.Listen("tcp", server_address)

	if err != nil {
		fmt.Println("Error listening on ", server_address)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	// send off a list of names that are currently in the server
	defer func() {
		//handleDisconnect(conn)
		conn.Close()
	}()

	// Add the user to the current list

	buf := make([]byte, 1<<19)

	// How do I send the initial user name?

	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading to buffer")
	}

	username := string(buf[:n])
	new_user := User{Name: username, Address: conn.RemoteAddr()}

	active_users = append(active_users, new_user)

	fmt.Println("Handling connection with " + username)
}

// When they join we need to add them to the user list
// When they lose we need to remove them from the user list

func handleDisconnect(u User) {
	// find that cocksucker in the big boy list
}
