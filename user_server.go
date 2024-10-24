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

func newHandleConnection(conn net.Conn) {
	buf := make([]byte, 1<<19)
	new_user := User{}

	defer func() {
		handleDisconnect(new_user)
		conn.Close()
	}()

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			fmt.Println("Connection closed by user:", conn.RemoteAddr())
			return
		}

		if new_user == (User{}) {
			createUser(conn)
		}
		fmt.Printf("Received: %s\n", string(buf[:n]))

		// create packet object
		// routeTraffic(packet)
		// if routeTraffic(packet) != nil { }
	}
}

func createUser(conn net.Conn) (user User) {
	// Fix this
	/* username := string(buf[:n])
	new_user := User{Name: username, Address: conn.RemoteAddr()}

	active_users = append(active_users, new_user)

	fmt.Println("User connected with the following data:")
	fmt.Println(new_user)

	sendUserList(conn) */
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

	fmt.Println("User connected with the following data:")
	fmt.Println(new_user)

	sendUserList(conn)

	defer func() {
		handleDisconnect(new_user)
		conn.Close()
	}()

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

	}
}

func sendUserList(conn net.Conn) {
	var message string = "Current users:"

	for i := 0; i < len(active_users); i++ {
		message += fmt.Sprintf("\n%d %s", i+1, active_users[i].Name)
	}

	conn.Write([]byte(message))
}

func handleNewConnection() (user User) {
	// Change the way we process this shit, need to pull from the packet object
	/* 	username := string(buf[:n])
	   	new_user := User{Name: username, Address: conn.RemoteAddr()}

	   	active_users = append(active_users, new_user)

	   	fmt.Println("User connected with the following data:")
	   	fmt.Println(new_user)

	   	sendUserList(conn) */
}

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

func handleBadRequest() {
	fmt.Println("Malformed request, something broke - check serialization / deserialization")
}
