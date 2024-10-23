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

func (u User) String() string {
	return fmt.Sprintf("Name: %s, Address: %s", u.Name, u.Address.String())
}

var active_users []User

func startServer() {
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

	buf := make([]byte, 1<<19)

	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading to buffer")
	}

	username := string(buf[:n])
	new_user := User{Name: username, Address: conn.RemoteAddr()}

	active_users = append(active_users, new_user)

	// fmt.Println("Handling connection with " + username)

	fmt.Println("User connected with the following data:")
	fmt.Println(new_user)

	// Send list of active users and hande input

	sendUserList(conn)

	defer func() {
		handleDisconnect(new_user)
		conn.Close()
	}()

	// I need to be able to handle state here - I need to know what kind of menu tree we're in
	// OR I just start building the GUI here and say fuck it

	// Making an API to facilitate those connections

	// P2P necessarily prevents me from hiding the IPs of the two fellas chatting
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			fmt.Println("Connection closed by user:", conn.RemoteAddr())
			return
		}

		fmt.Printf("Received: %s\n", string(buf[:n]))
		// Okay, now move all the route traffic guys into the router and go from there
		// Need to figure out how to convert the traffice back into a packet object
		// Serializing and deserializing JSON
		//routeTraffic((string(buf[:n])))
	}
}

/* func sendUserList(conn net.Conn) {
	var message string = "Current users:"

	for i := 0; i < len(active_users); i++ {
		message += fmt.Sprintf("\n%d %s", i+1, active_users[i].Name)
	}

	conn.Write([]byte(message))
} */

func handleDisconnect(u User) {
	var new_active_users []User

	for i := 0; i < len(active_users); i++ {
		if u != active_users[i] {
			new_active_users = append(new_active_users, u)
		}
	}

	active_users = new_active_users

	fmt.Println("Active users after disconnect: ", active_users)

}

/*
func askPermission(asker User, recipient User) {
	// Send a message to the recipient saying asker wants to chat with them

	// Accept or decline

	// Setup a connection between the two of them

}
*/
