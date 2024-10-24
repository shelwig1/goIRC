package main

import (
	"net"
)

type Packet struct {
	Goal string
	Info string
}

func routeTraffic(packet Packet, conn net.Conn) {

	switch packet.Goal {
	/* 	case "initial_connection":
	//initialConnection(packet) */
	case "request_list":
		sendUserList(conn)
	case "request_user":
		sendUserRequest()
	case "request_user_response":
		processResponse()
	default:
		handleBadRequest()
	}

}

/*
func sendUserList(conn net.Conn) {
	var message string = "Current users:"

	for i := 0; i < len(active_users); i++ {
		message += fmt.Sprintf("\n%d %s", i+1, active_users[i].Name)
	}

	conn.Write([]byte(message))
	//fmt.Println(active_users)
}

func processResponse() {

}

func handleBadRequest() {

}
*/
